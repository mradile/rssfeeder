package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func (c *client) makeHTTPRequest(method, url string, data interface{}) (*http.Request, error) {
	byts, err := toJSONBytes(data)
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal struct to json")
	}
	uri := fmt.Sprintf("%s%s", c.cfg.Hostname, url)

	if c.debug {
		fmt.Printf("creating request: [%s] [%s]\n", method, uri)
	}

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

func (c *client) setAuthHeader(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.cfg.Token)
}
