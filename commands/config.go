package config

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
)

type berks struct {
	Main     main
	Cookbook []cookbook
}

type main struct {
	Library    string
	Definition string
	Host       string
}

type cookbook struct {
	Path string
	Host string
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
path = "engineyard/ey-cloud-recipes/cookbooks/env_vars"
`
	f.Write([]byte(defaultFormat))
	f.Close()

	return nil
}

func Parse(path string) berks {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var berks berks
	if _, err := toml.Decode(string(dat), &berks); err != nil {
	}
	return berks
}
