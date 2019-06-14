package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mradile/rssfeeder/pkg/rest"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Login(c echo.Context) error {
	var logReq rest.LoginRequest
	if err := c.Bind(&logReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request body invalid")
	}

	user, err := h.users.Get(logReq.Login)
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

	accessToken, err := h.getAccessToken(logReq.Login)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not authenticate")
	}
	refreshToken, err := h.getRefreshToken(logReq.Login)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not authenticate")
	}

	return c.JSONPretty(http.StatusOK, &rest.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, "  ")
}

func (h *Handler) RefreshAccessToken(c echo.Context) error {
	var refreshReq rest.RefreshAccessTokenRequest
	if err := c.Bind(&refreshReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request body invalid")
	}

	token, err := jwt.Parse(refreshReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(h.cfg.SessionSecret), nil
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not renew access token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "could not renew access token")
	}
	if !token.Valid {
		return echo.NewHTTPError(http.StatusBadRequest, "refresh token invalid")
	}

	login, ok := claims["login"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "refresh token invalid")
	}

	if user, err := h.users.Get(login); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not renew access token")
	} else if user == nil {
		return echo.NewHTTPError(http.StatusForbidden, "login not found")
	}

	aToken, err := h.getAccessToken(login)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not renew access token")
	}

	refreshRes := &rest.RefreshAccessTokenResponse{
		AccessToken: aToken,
	}

	return c.JSONPretty(http.StatusOK, refreshRes, "  ")
}

func (h *Handler) getAccessToken(login string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["login"] = login
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	claims["iat"] = time.Now().Unix()

	t, err := token.SignedString([]byte(h.cfg.SessionSecret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (h *Handler) getRefreshToken(login string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["login"] = login
	claims["exp"] = time.Now().Add(h.cfg.SessionTTL).Unix()
	claims["iat"] = time.Now().Unix()

	t, err := token.SignedString([]byte(h.cfg.SessionSecret))
	if err != nil {
		return "", err
	}
	return t, nil
}
