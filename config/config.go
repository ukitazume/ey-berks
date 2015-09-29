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
	WorkingDirName = ".ey-berks"
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
	host string
	Repo string
	Name string
}

func (c *Cookbook) WorkingRootPath() string {
	return filepath.Join(os.Getenv("HOME"), WorkingDirName)
}

func (c *Cookbook) WorkingRepoPath() string {
	return filepath.Join(c.WorkingRootPath(), c.Host(), c.Repo)
}

func (c *Cookbook) WorkingPath() string {
	return filepath.Join(c.WorkingRepoPath(), c.Path)
}

func (c *Cookbook) Host() string {
	if c.host != "" {
		return c.host
	} else {
		return "github.com"
	}
}

func (c *Cookbook) RemoteUrl() string {
	return "git://" + c.Host() + "/" + c.Repo
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
