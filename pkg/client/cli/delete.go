package cli

import (
	"fmt"
	"github.com/mradile/rssfeeder/pkg/client/configuration"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"net/http"
	"strconv"
)

func (c *client) DeleteFeedEntry() cli.Command {
	cmd := cli.Command{
		Name:      "delete-entry",
		ShortName: "de",
		Usage:     "delete a feed entry",
		UsageText: "de 23",
		Action: func(c *cli.Context) error {
			setLogLevel(c)
			cfg, err := configuration.Load(c.GlobalString("cfg"))
			if err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not add entry"), 1)
			}
			Client.cfg = cfg

			id, err := getIdFromArgs(c)
			if err != nil {
				return cli.NewExitError(err, 1)
			}

			url := fmt.Sprintf("/api/v1/entry/%d", id)
			err = deleteRequest(id, url)
			if err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not delete entry"), 1)
			}

			LogInfo(fmt.Sprintf("deleted entry [%d] \n", id))

			return nil
		},
	}
	return cmd
}

func (c *client) DeleteFeed() cli.Command {
	cmd := cli.Command{
		Name:      "delete-feed",
		ShortName: "df",
		Usage:     "delete a feed ",
		UsageText: "def 42",
		Action: func(c *cli.Context) error {
			setLogLevel(c)
			cfg, err := configuration.Load(c.GlobalString("cfg"))
			if err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not add entry"), 1)
			}
			Client.cfg = cfg

			id, err := getIdFromArgs(c)
			if err != nil {
				return cli.NewExitError(err, 1)
			}

			url := fmt.Sprintf("/api/v1/feed/%d", id)
			err = deleteRequest(id, url)
			if err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not delete feed"), 1)
			}

			LogInfo(fmt.Sprintf("deleted feed [%d] \n", id))

			return nil
		},
	}
	return cmd
}

func getIdFromArgs(c *cli.Context) (int, error) {
	idStr := c.Args().Get(0)
	if idStr == "" {
		return 0, fmt.Errorf("please specify an id")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("please specify a valid id")
	}
	return id, nil
}

func deleteRequest(id int, url string) error {
	req, err := Client.makeHTTPRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	err = Client.setAuthHeader(req)
	if err != nil {
		return err
	}

	return Client.getResponse(req, nil, http.StatusNoContent)
}
