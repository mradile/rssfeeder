package cmd

import (
	"fmt"
	"github.com/mradile/rssfeeder/pkg/client/configuration"
	"github.com/mradile/rssfeeder/pkg/rest"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"net/http"
)

func (c *client) AddEntry() cli.Command {
	cmd := cli.Command{
		Name:      "add",
		ShortName: "a",
		Usage:     "add an URI to a feed",
		UsageText: "add http://example.org/item news",
		Action: func(c *cli.Context) error {
			setLogLevel(c)
			cfg, err := configuration.Load(c.GlobalString("cfg"))
			if err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not add entry"), 1)
			}
			Client.cfg = cfg

			entry := c.Args().Get(0)
			if entry == "" {
				return cli.NewExitError("please specify an uri", 1)
			}

			category := c.Args().Get(1)
			if category == "" {
				category = ""
			}

			addReq := &rest.AddEntryRequest{
				URI:      entry,
				Category: category,
			}

			req, err := Client.makeHTTPRequest("POST", "/api/v1/entry", addReq)
			if err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not add entry"), 1)
			}
			Client.setAuthHeader(req)

			var addRes rest.AddEntryResponse
			if err := Client.getResponse(req, &addRes, http.StatusOK); err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not add entry"), 1)
			}

			fmt.Printf("added entry [%s] in category [%s] with id [%d]\n", entry, addRes.Category, addRes.ID)

			return nil
		},
		OnUsageError:       nil,
		Subcommands:        nil,
		Flags:              nil,
		SkipFlagParsing:    false,
		SkipArgReorder:     false,
		HideHelp:           false,
		Hidden:             false,
		HelpName:           "",
		CustomHelpTemplate: "",
	}
	return cmd
}
