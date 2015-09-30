package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

const (
	ConfigFileName = "EyBerksfile"
	CacheDirName   = ".ey-berks"
)

type CodeResourceOperator interface {
	RemoteHost() string
	RemotoRepoUrl() string   /* git@github.com:engineyard/ey-cloud-recipes */
	CacheRepoPath() string   /* /home/deploy/.ey-berks/github.com/engineyard/ey-cloud-recipes */
	DesticationPath() string /* env_vars */
	SourcePath() string      /* env_vars */
}

func (c *CodeResource) RecipeName() string {
	s := strings.Split(c.DesticationPath(), "/")
	name := s[len(s)-1]
	return name
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
	return filepath.Join(os.Getenv("HOME"), CacheDirName, c.RemoteHost(), c.Repo)
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
	if c.Host == "bitbucket.org" {
		return "https://" + c.RemoteHost() + "/" + c.Repo + ".git"
	} else {
		return "git://" + c.RemoteHost() + "/" + c.Repo
	}
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
repo = "engineyard/ey-cloud-recipes"
path = "cookbooks/env_vars"

[[cookbook]]
host = "bitbucket.org"
repo = "ukitazume/ey-mini-cookbooks"
path = "cookbooks/custom_nginx"
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
