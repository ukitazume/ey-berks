package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const (
	ConfigFileName = "EYBerksfile"
)

type Berks struct {
	Main     Main
	Cookbook []Cookbook
}

type Main struct {
	Library    string
	Definition string
	Host       string
}

type Cookbook struct {
	Path string
	Host string
	Repo string
	NAME string
}

func Create(path string) error {
	fullPath := filepath.Join(path, ConfigFileName)

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defaultFormat := `[main]
library = "engineyard/ey-cloud-recipes/main/libraries"
definition = "engineyard/ey-cloud-recipes/main/definitions"

[[cookbook]]
repo = "engineyard/ey-cloud-recipes"
path = "cookbooks/env_vars"
`
	f.Write([]byte(defaultFormat))
	f.Close()

	return nil
}

func Parse(path string) Berks {
	fullPath := filepath.Join(path, ConfigFileName)
	dat, err := ioutil.ReadFile(fullPath)
	if err != nil {
		log.Fatal(err)
	}
	var berks Berks
	if _, err := toml.Decode(string(dat), &berks); err != nil {
		log.Fatal(err)
	}
	return berks
}
