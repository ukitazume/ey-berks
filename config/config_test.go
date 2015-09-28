package config

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func removeTmpFile(t *testing.T) {
	err := os.Remove("/tmp/EYBerksfile")
	if err != nil {
		t.Errorf("failed to remove /tmp/EYBerksfile")
	}
}

func TestCreate(t *testing.T) {
	Create("/tmp")
	dat, err := ioutil.ReadFile("/tmp/EYBerksfile")
	if err != nil {
		t.Errorf("get error when opening EYBerksfile")
	}
	actual := string(dat)
	r, _ := regexp.Compile("[main]")
	if !r.MatchString(actual) {
		t.Errorf("not include %v in %v", r, actual)
	}
	removeTmpFile(t)

}

func TestParse(t *testing.T) {
	Create("/tmp")
	berks := Parse("/tmp")
	if berks.Main.Library != "engineyard/ey-cloud-recipes/main/libraries" {
		t.Errorf("not match")
	}
	if berks.Cookbook[0].Repo != "engineyard/ey-cloud-recipes" {
		t.Errorf("not match")
	}
	if berks.Cookbook[0].Path != "cookbooks/env_vars" {
		t.Errorf("not match")
	}

	removeTmpFile(t)
}
