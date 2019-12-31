package configuration

import "time"

type Configuration struct {
	Version       string
	DBPath        string
	Addr          string
	Hostname      string
	SessionSecret string
	SessionTTL    time.Duration
}
