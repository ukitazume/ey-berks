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
	ConfigFileName  = "EyBerksfile"
	CacheDirName    = ".ey-berks"
	CookBookDirName = "cookbooks"
)

type Options struct {
	ConfigFileName string
	TargetDirName  string
}

func DefaultOption() Options {
	return Options{ConfigFileName, CookBookDirName}
}

type CodeResourceOperator interface {
	RemoteHost() string
	RemotoRepoUrl() string   /* git@github.com:engineyard/ey-cloud-recipes */
	CacheRepoPath() string   /* /home/deploy/.ey-berks/github.com/engineyard/ey-cloud-recipes */
	DesticationPath() string /* cookbooks/env_vars */
	RecipeName() string
	SourcePath() string
}

func (c *CodeResource) RecipeName() string {
	s := strings.Split(c.Path, "/")
	name := s[len(s)-1]
	return name
}

func (c *CodeResource) DesticationPath() string {
	if c.Name == "" {
		return filepath.Join(CookBookDirName, c.RecipeName())
	} else {
		return filepath.Join(CookBookDirName, c.Name)
	}
}

func (c *CodeResource) SourcePath() string {
	return c.Path
}

func (c *CodeResource) CacheRepoPath() string {
	return filepath.Join(os.Getenv("HOME"), CacheDirName, c.RemoteHost(), c.Repo)
}

type CodeResource struct {
	Path string
	Name string
	Host string
	Repo string
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

func IsExistConfigFile(path string, opts Options) bool {
	fullPath := filepath.Join(path, opts.ConfigFileName)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func Create(path string, opts Options) error {
	fullPath := filepath.Join(path, opts.ConfigFileName)

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defaultFormat := `[library]
repo = "engineyard/ey-cloud-recipes"
path = "cookbooks/main/libraries"
name = "main/libraries"

[definition]
repo = "engineyard/ey-cloud-recipes"
path = "cookbooks/main/definitions"
name = "main/definitions"

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

func Parse(path string, options Options) Berks {
	fullPath := filepath.Join(path, options.ConfigFileName)
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
