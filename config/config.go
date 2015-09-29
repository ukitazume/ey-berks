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

type CodeResourceOperator interface {
	RemoteHost() string
	RemotoRepoUrl() string   /* git://github.com/egnineyard/ey-cloud-recipes */
	CacheRepoPath() string   /* /home/deploy/.ey-berks/github.com/engineyard/ey-cloud-recipes */
	DesticationPath() string /* env_vars */
	SourcePath() string      /* env_vars */
}

func (c *CodeResource) DesticationPath() string {
	if c.DistPath == "" {
		return c.Path
	} else {
		return c.DistPath
	}
}

func (c *CodeResource) SourcePath() string {
	if c.SrcPath == "" {
		return c.Path
	} else {
		return c.SrcPath
	}
}

func (c *CodeResource) CacheRepoPath() string {
	return filepath.Join(os.Getenv("HOME"), WorkingDirName, c.RemoteHost(), c.Repo)
}

type CodeResource struct {
	Path     string
	SrcPath  string
	DistPath string
	Host     string
	Repo     string
	Name     string
	CodeResourceOperator
}

type Berks struct {
	Library    Library
	Definition Definition
	Cookbooks  []Cookbook `toml:"cookbook"`
}

type Library struct {
	*CodeResource
}

type Definition struct {
	*CodeResource
}

type Cookbook struct {
	*CodeResource
}

func (c *CodeResource) RemoteHost() string {
	if c.Host != "" {
		return c.Host
	} else {
		return "github.com"
	}
}

func (c *CodeResource) RemotoRepoUrl() string {
	return "git://" + c.RemoteHost() + "/" + c.Repo
}

func Create(path string) error {
	fullPath := filepath.Join(path, ConfigFileName)

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defaultFormat := `[library]
repo = "engineyard/ey-cloud-recipes"
path = "cookbooks/main/libraries"

[definition]
repo = "engineyard/ey-cloud-recipes"
path = "cookbooks/main/definitions"

[[cookbook]]
name = "env"
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
