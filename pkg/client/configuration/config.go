package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"path"
)

const ConfigFileName = "rssfeeder.json"

type Configuration struct {
	Hostname string `json:"hostname"`
	Token    string `json:"token"`
	Login    string `json:"login"`
}

func Save(cfg *Configuration, cfgPath string) error {
	byts := ToJSON(cfg)
	filepath := path.Join(cfgPath, ConfigFileName)

	err := ioutil.WriteFile(filepath, byts, 0600)
	if err != nil {
		return errors.Wrapf(err, "could write configuration file [%s]", filepath)
	}
	fmt.Printf("config written to [%s]\n", filepath)

	return nil
}

func Load(cfgPath string) (*Configuration, error) {
	filepath := path.Join(cfgPath, ConfigFileName)
	byts, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrapf(err, "could not open configuration file [%s]", filepath)
	}
	var cfg Configuration
	if err := json.Unmarshal(byts, &cfg); err != nil {
		return nil, errors.Wrapf(err, "could not parse configuration file [%s]", filepath)
	}
	return &cfg, nil
}

func ToJSON(cfg *Configuration) []byte {
	byts, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("could not marshal config to json: %s", err))
	}
	return byts
}
