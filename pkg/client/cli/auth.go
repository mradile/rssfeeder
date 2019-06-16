package cli

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mradile/rssfeeder/pkg/client/configuration"
	"github.com/mradile/rssfeeder/pkg/rest"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

func (c *client) setAuthHeader(req *http.Request) error {
	aToken, err := c.checkAccessToken()
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+aToken)
	return nil
}

func (c *client) checkAccessToken() (string, error) {
	aToken := c.cfg.AccessToken

	parser := new(jwt.Parser)
	claims := jwt.MapClaims{}
	if _, _, err := parser.ParseUnverified(aToken, claims); err != nil {
		return "", errors.Wrap(err, "could not parse token")
	}

	expF, ok := claims["exp"].(float64)
	if !ok {
		return "", errors.New("could not parse exp field to float64")
	}

	exp := int64(expF)
	nowT := time.Now()
	now := nowT.Add(time.Second * 100).Unix()
	if now >= exp {
		LogDebug("new access token needed")
		err := c.refreshAccessToken()
		if err != nil {
			return "", errors.Wrap(err, "retrieving new access token failed")
		}
	}

	return c.cfg.AccessToken, nil
}

func (c *client) refreshAccessToken() error {
	reqATR := &rest.RefreshAccessTokenRequest{
		RefreshToken: c.cfg.RefreshToken,
		Login:        c.cfg.Login,
	}
	req, err := c.makeHTTPRequest("POST", "/auth/refresh", reqATR)
	if err != nil {
		return err
	}

	var resATR rest.RefreshAccessTokenResponse
	err = c.getResponse(req, &resATR, http.StatusOK)
	if err != nil {
		return err
	}

	if resATR.AccessToken == "" {
		return errors.New("received access token is empty")
	}

	c.cfg.AccessToken = resATR.AccessToken
	err = configuration.Save(c.cfg, "")
	if err != nil {
		return errors.New("could not save new access token")
	}

	return nil
}
