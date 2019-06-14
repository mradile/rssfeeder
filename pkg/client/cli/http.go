package cli

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mradile/rssfeeder/pkg/client/configuration"
	"github.com/mradile/rssfeeder/pkg/rest"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

func (c *client) makeHTTPRequest(method, url string, data interface{}) (*http.Request, error) {
	byts, err := toJSONBytes(data)
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal struct to json")
	}
	uri := fmt.Sprintf("%s%s", c.cfg.Hostname, url)

	LogDebug(fmt.Sprintf("creating request for %s - %s", method, uri))

	req, err := http.NewRequest(method, uri, byts)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *client) getResponse(req *http.Request, response interface{}, httpCode int) error {
	res, err := Client.http.Do(req)
	if err != nil {
		return err
	}

	//noinspection GoUnhandledErrorResult
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	code := res.StatusCode
	if code != httpCode {
		return fmt.Errorf("http code was [%d], expected [%d]: %s", code, httpCode, string(data))
	}

	if err := json.Unmarshal(data, response); err != nil {
		return fmt.Errorf("could not unmarshal response '%s': %v", string(data), err)
	}

	return nil
}

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
	now := time.Now().Add(time.Second * 100).Unix()
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
