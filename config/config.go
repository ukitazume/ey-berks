package config

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
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
	var fullPath bytes.Buffer
	fullPath.WriteString(path)
	fullPath.WriteString("/EYBerksfile")

	f, err := os.Create(fullPath.String())
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
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var berks Berks
	if _, err := toml.Decode(string(dat), &berks); err != nil {
	}
	return berks
}
