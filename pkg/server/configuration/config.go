package configuration

import "time"

type Configuration struct {
	DBPath        string
	Addr          string
	Hostname      string
	SessionSecret string
	SessionTTL    time.Duration
}
