package http

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mradile/rssfeeder"
	"github.com/mradile/rssfeeder/pkg/server/adding"
	"github.com/mradile/rssfeeder/pkg/server/configuration"
	"github.com/mradile/rssfeeder/pkg/server/deleting"
	"github.com/mradile/rssfeeder/pkg/server/viewing"
)

type Server struct {
	e          *echo.Echo
	addr       string
	cfg        *configuration.Configuration
	httpServer *http.Server
	users      rssfeeder.UserStorage
	adder      adding.AddingService
	deleter    deleting.DeletingService
	viewer     viewing.ViewingService
}

type Services struct {
}

func NewServer(cfg *configuration.Configuration, users rssfeeder.UserStorage, adder adding.AddingService, deleter deleting.DeletingService, viewer viewing.ViewingService) *Server {
	e := echo.New()

	//recover from panics
	e.Use(middleware.Recover())

	//add a unique id to each request
	e.Use(middleware.RequestID())

	//set max body size for requests
	e.Use(middleware.BodyLimit("10KB"))

	e.HideBanner = true
	e.HidePort = true

	s := &Server{
		e:       e,
		addr:    cfg.Addr,
		cfg:     cfg,
		users:   users,
		adder:   adder,
		deleter: deleter,
		viewer:  viewer,
		httpServer: &http.Server{
			Addr:              cfg.Addr,
			ReadTimeout:       60 * time.Second,  // time to read request from client
			ReadHeaderTimeout: 10 * time.Second,  // time to read header, low value to cope with malicious behavior
			WriteTimeout:      20 * time.Second,  // time write to the client
			IdleTimeout:       120 * time.Second, // time between keep-alives requests before connection is closed
		},
	}

	auth := e.Group("/auth")
	auth.POST("/login", s.Login)

	//http://localhost:3000/feeds/mre/fla/.rss
	rss := e.Group("/feeds")
	rss.GET("/:login/:category/:type/.rss", s.RSSFeed)

	api := e.Group("/api/v1")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(cfg.SessionSecret),
		ErrorHandler: func(err error) error {
			if e, ok := err.(*jwt.ValidationError); ok {
				if e.Errors == jwt.ValidationErrorExpired {
					return echo.NewHTTPError(http.StatusUnauthorized, "expired")
				}
			}
			return err
		},
	}))
	api.POST("/entry", s.AddEntry)
	api.DELETE("/entry/:id", s.DeleteEntry)
	api.GET("/feed", s.ListFeeds)

	return s
}

func (s *Server) Start() error {
	return s.e.StartServer(s.httpServer)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}

func getLoginFromContext(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	login, ok := claims["login"].(string)
	if !ok {
		return ""
	}
	return login
}
