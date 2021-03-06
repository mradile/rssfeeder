package cli

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"

	"github.com/mradile/rssfeeder/pkg/client/configuration"
	"github.com/mradile/rssfeeder/pkg/rest"

	"github.com/urfave/cli"
)

func (c *client) Login() cli.Command {

	cmd := cli.Command{
		Name:  "login",
		Usage: "this command makes a login for given user and writes a config file",
		Action: func(c *cli.Context) error {
			setLogLevel(c)

			hostname := c.String("hostname")
			if hostname == "" {
				return cli.NewExitError("you must provide a valid hostname", 1)
			}
			Client.cfg = &configuration.Configuration{
				Hostname: hostname,
			}

			//ToDo read password from term when omitted

			loginReq := &rest.LoginRequest{
				Login:    c.String("login"),
				Password: c.String("password"),
			}
			req, err := Client.makeHTTPRequest("POST", "/auth/login", loginReq)
			if err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not login"), 1)
			}

			var token rest.LoginResponse
			if err := Client.getResponse(req, &token, http.StatusOK); err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not login"), 1)
			}

			if token.AccessToken == "" || token.RefreshToken == "" {
				return cli.NewExitError("could not login, received empty token", 1)
			}

			LogInfo(fmt.Sprintf("successfully logged in as user %s", loginReq.Login))

			Client.cfg.AccessToken = token.AccessToken
			Client.cfg.RefreshToken = token.RefreshToken
			Client.cfg.Login = loginReq.Login
			if err := configuration.Save(Client.cfg, c.GlobalString("cfg")); err != nil {
				return cli.NewExitError(errors.Wrap(err, "could not save configuration after login"), 1)
			}

			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "login, l",
				Usage:  "`username`",
				EnvVar: "LOGIN",
				Value:  "",
			},
			cli.StringFlag{
				Name:   "password, p",
				Usage:  "`password`",
				EnvVar: "PASSWORD",
				Value:  "",
			},
			cli.StringFlag{
				Name:   "hostname",
				Usage:  "`hostname`",
				EnvVar: "HOSTNAME",
				Value:  "http://localhost:3000",
			},
		},
	}
	return cmd
}
