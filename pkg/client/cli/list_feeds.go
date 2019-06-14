package cli

import (
	"fmt"
	"github.com/mradile/rssfeeder/pkg/client/configuration"
	"github.com/mradile/rssfeeder/pkg/rest"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"net/http"
)

func (c *client) ListFeeds() cli.Command {
	cmd := cli.Command{
		Name:      "list",
		ShortName: "l",
		Usage:     "list all feeds with their URIs",
		UsageText: "list",
		Action: func(c *cli.Context) error {
			setLogLevel(c)
			cfg, err := configuration.Load(c.GlobalString("cfg"))
			if err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not list feeds"), 1)
			}
			Client.cfg = cfg

			req, err := Client.makeHTTPRequest("GET", "/api/v1/feed", nil)
			if err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not list feeds"), 1)
			}
			err = Client.setAuthHeader(req)
			if err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not list feeds"), 1)
			}

			var list rest.FeedListResponse
			if err := Client.getResponse(req, &list, http.StatusOK); err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not list feeds"), 1)
			}

			fmt.Println("Feeds:")
			for _, feed := range list.Feeds {
				fmt.Printf("%s\n", feed.Name)
				for _, uri := range feed.URIs {
					fmt.Printf("\t%s\n", uri)
				}
			}

			return nil
		},
	}
	return cmd
}
