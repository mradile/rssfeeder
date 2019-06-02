package cli

import (
	"bytes"
	"encoding/json"
	"github.com/mradile/rssfeeder/pkg/client/configuration"
	"github.com/urfave/cli"
	"io"
	"net/http"
	"time"
)

type client struct {
	http    *http.Client
	cfg     *configuration.Configuration
	debug   bool
	verbose bool
}

var Client = &client{
	http: &http.Client{
		Timeout: time.Duration(time.Second * 60),
	},
}

func toJSONBytes(data interface{}) (io.Reader, error) {
	byts, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(byts), nil
}

func setLogLevel(c *cli.Context) {
	if c.GlobalBool("debug") {
		Client.debug = true
		return
	}
	if c.GlobalBool("verbose") {
		Client.verbose = true
	}
}
