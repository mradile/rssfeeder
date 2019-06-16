package cli

import (
	"encoding/json"
	"fmt"
	"github.com/mradile/rssfeeder/pkg/rest"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

func (c *client) makeHTTPRequest(method, url string, data interface{}) (*http.Request, error) {

	var reader io.Reader
	if data != nil {
		var err error
		reader, err = toJSONReader(data)
		if err != nil {
			return nil, errors.Wrap(err, "could not marshal struct to json")
		}
	}

	uri := fmt.Sprintf("%s%s", c.cfg.Hostname, url)

	LogDebug(fmt.Sprintf("creating request for %s - %s", method, uri))

	req, err := http.NewRequest(method, uri, reader)
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

	code := res.StatusCode
	//if 204 is expected and response is 204 there is no need for reading the body
	if httpCode == http.StatusNoContent && code == http.StatusNoContent {
		return nil
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if code != httpCode {
		var er rest.ErrorResponse
		err := json.Unmarshal(data, &er)

		if err == nil {
			if c.debug {
				return fmt.Errorf("http code was [%d], expected [%d]: %s", code, httpCode, er.Message)
			}
			return fmt.Errorf(er.Message)
		}
		return fmt.Errorf("http code was [%d], expected [%d]: %s", code, httpCode, string(data))
	}

	if err := json.Unmarshal(data, response); err != nil {
		return fmt.Errorf("could not unmarshal response '%s': %v", string(data), err)
	}

	return nil
}
