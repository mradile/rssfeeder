package http

import (
	"net/http"
	"time"

	"github.com/mradile/rssfeeder/pkg/rest"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) Login(c echo.Context) error {

	var logReq rest.LoginRequest
	if err := c.Bind(&logReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request body invalid")
	}

	user, err := s.users.Get(logReq.Login)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not authenticate")
	}
	if user == nil {
		return echo.NewHTTPError(http.StatusForbidden, "user and / or password are incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logReq.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "user and / or password are incorrect")
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	ttl := time.Now().Add(s.cfg.SessionTTL)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["login"] = logReq.Login
	claims["exp"] = ttl.Unix()
	claims["iat"] = time.Now().Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(s.cfg.SessionSecret))
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, &rest.LoginResponse{Token: t}, "  ")
}
