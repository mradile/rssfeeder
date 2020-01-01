package cli

import (
	"bufio"
	"fmt"
	"github.com/mradile/rssfeeder/pkg/client/configuration"
	"github.com/mradile/rssfeeder/pkg/rest"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"net/http"
	"os"
)

func (c *client) AddEntry() cli.Command {
	cmd := cli.Command{
		Name:      "add",
		ShortName: "a",
		Usage:     "add an URI to a feed",
		UsageText: `add http://example.org/item news
   add - news < file_with_uris.txt`,
		Description: `Either specify the uri directy as first parameter after the 'add' command or use the '-' for 
   reading from stdin.`,
		Action: func(c *cli.Context) error {
			setLogLevel(c)
			cfg, err := configuration.Load(c.GlobalString("cfg"))
			if err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not add entry"), 1)
			}
			Client.cfg = cfg

			entry := c.Args().Get(0)
			if entry == "" {
				return cli.NewExitError("please specify an uri or use - for reading from stdin", 1)
			}

			feedName := c.Args().Get(1)

			if entry == "-" {
				if err := readFromStdin(feedName); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}

			} else {
				addReq := &rest.AddEntryRequest{
					URI:      entry,
					FeedName: feedName,
				}
				if err := addFeedEntry(addReq); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
			}

			return nil
		},
	}
	return cmd
}

func readFromStdin(feedName string) error {
	hasError := false
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		entry := scanner.Text()
		if len(entry) != 0 {
			addReq := &rest.AddEntryRequest{
				URI:      entry,
				FeedName: feedName,
			}
			if err := addFeedEntry(addReq); err != nil {
				fmt.Println(err.Error())
				hasError = true
			}
		} else {
			break
		}
	}
	if hasError {
		return errors.New("at least on error occured")
	}
	return nil
}

func addFeedEntry(addReq *rest.AddEntryRequest) error {
	req, err := Client.makeHTTPRequest("POST", "/api/v1/entry", addReq)
	if err != nil {
		return errors.Wrap(err, "could not add entry")
	}
	err = Client.setAuthHeader(req)
	if err != nil {
		return errors.Wrap(err, "could not add entry")
	}

	var addRes rest.AddEntryResponse
	if err := Client.getResponse(req, &addRes, http.StatusOK); err != nil {
		return errors.Wrap(err, "could not add entry")
	}

	LogInfo(fmt.Sprintf("added entry [%s] in feed [%s] with id [%d]\n", addReq.URI, addRes.FeedName, addRes.ID))
	return nil
}
