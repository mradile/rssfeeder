package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mradile/rssfeeder/pkg/server/viewing"

	"github.com/mradile/rssfeeder/pkg/server/deleting"

	"github.com/mradile/rssfeeder/pkg/server/adding"

	"github.com/mradile/rssfeeder/pkg/server/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/sirupsen/logrus"

	"github.com/mradile/rssfeeder"

	"github.com/mradile/rssfeeder/pkg/server/storage"

	"github.com/mradile/rssfeeder/pkg/server/configuration"

	"github.com/urfave/cli"
)

var version = "dev-snapshot"

var debug bool
var verbose bool

func main() {

	cliApp := cli.NewApp()
	cliApp.Name = "rssfeeder server"
	cliApp.Usage = ""
	cliApp.Version = version

	//host=%s port=%d user=%s dbname=%s password=%s sslmode=%s

	cliApp.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "prints debug level log messages",
			EnvVar:      "LOG_DEBUG",
			Destination: &debug,
		},
		cli.BoolFlag{
			Name:        "verbose",
			Usage:       "prints info level log messages",
			EnvVar:      "LOG_VERBOSE",
			Destination: &verbose,
		},
		cli.StringFlag{
			Name:   "db",
			Usage:  "`directory` where to store the database",
			EnvVar: "DB",
			Value:  "",
		},
		cli.IntFlag{
			Name:   "port",
			Usage:  "`port` to listen",
			EnvVar: "PORT",
			Value:  3000,
		},
		cli.StringFlag{
			Name:   "hostname",
			Usage:  "`hostname` that will be used in the rss feeds",
			EnvVar: "HOSTNAME",
			Value:  "http://localhost:3000",
		},
		cli.StringFlag{
			Name:   "secret",
			Usage:  "`secret` string for jwt session",
			EnvVar: "SECRET",
		},
		cli.DurationFlag{
			Name:   "session-ttl",
			Usage:  "`ttl` of jwt session",
			EnvVar: "SESSION_TTL",
			Value:  time.Duration(time.Hour * 24 * 30),
		},
		cli.BoolFlag{
			Name:   "create-user",
			Usage:  "create/updates a user in the database",
			EnvVar: "CREATE_USER",
		},
		cli.StringFlag{
			Name:   "login",
			Usage:  "`login` for creating/updating a user in the database",
			EnvVar: "LOGIN",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "`password` for creating/updating a user in the database",
			EnvVar: "PASSWORD",
		},
	}

	cliApp.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "start the server",
			Action: func(c *cli.Context) error {
				cfg := makeConfig(c)
				db, err := storage.NewStormDB(cfg)
				if err != nil {
					return err
				}

				user := storage.NewUserStorage(db)
				feedEntries := storage.NewFeedEntryStorage(db)
				adder := adding.NewAddingService(feedEntries)
				deleter := deleting.NewDeletingService(feedEntries)
				viewer := viewing.NewViewingService(feedEntries)

				if err := createUser(user, c); err != nil {
					return err
				}

				server := http.NewServer(cfg, user, adder, deleter, viewer)
				return server.Start()
			},
			Flags: []cli.Flag{},
		},
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		panic(fmt.Sprintf("could not initialize app: %s", err))
	}
}

func createUser(users rssfeeder.UserStorage, c *cli.Context) error {
	if !c.GlobalBool("create-user") {
		return nil
	}

	login := c.GlobalString("login")
	user, err := users.Get(login)
	if err != nil {
		return err
	}

	password := c.GlobalString("password")
	hashedPW, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if user == nil {
		logrus.WithFields(logrus.Fields{
			"login": login,
		}).Info("adding user")

		return users.Add(&rssfeeder.User{
			Login:    login,
			Password: string(hashedPW),
		})
	}

	logrus.WithFields(logrus.Fields{
		"login": login,
	}).Info("updating user")

	user.Login = login
	user.Password = string(hashedPW)

	return users.Update(user)
}

func makeConfig(c *cli.Context) *configuration.Configuration {
	cfg := &configuration.Configuration{
		DBPath:        c.GlobalString("db"),
		Addr:          ":" + strconv.Itoa(c.GlobalInt("port")),
		Hostname:      c.GlobalString("hostname"),
		SessionSecret: c.GlobalString("secret"),
		SessionTTL:    c.GlobalDuration("session-ttl"),
	}

	if cfg.SessionSecret == "" {
		panic("secret must not be empty")
	}

	return cfg
}
