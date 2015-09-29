package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func removeTmpFile(t *testing.T) {
	err := os.Remove("/tmp/EYBerksfile")
	if err != nil {
		t.Errorf("failed to remove /tmp/EYBerksfile")
	}
}

func createConfig(t *testing.T) string {
	Create("/tmp")
	dat, err := ioutil.ReadFile("/tmp/EYBerksfile")
	if err != nil {
		t.Errorf("get error when opening EYBerksfile")
	}
	return string(dat)
}

func TestCookbookRemoteRepoUrl(t *testing.T) {
	createConfig(t)
	berks := Parse("/tmp")

	actuals := [3]string{
		berks.Library.RemotoRepoUrl(),
		berks.Definition.RemotoRepoUrl(),
		berks.Cookbooks[0].RemotoRepoUrl(),
	}
	expect := "git://github.com/engineyard/ey-cloud-recipes"

	for _, actual := range actuals {
		if expect != actual {
			t.Errorf("wrong RemoteRepoURL expect %v, got %v", expect, actual)
		}
	}

	removeTmpFile(t)
}

func TestCacheRepoPath(t *testing.T) {
	createConfig(t)
	berks := Parse("/tmp")

	actual := berks.Library.CacheRepoPath()
	expect := filepath.Join(os.Getenv("HOME"), WorkingDirName, berks.Library.RemoteHost(), berks.Library.Repo)
	if expect != actual {
		t.Errorf("wrong RemoteRepoURL expect %v, got %v", expect, actual)
	}

	removeTmpFile(t)
}

func TestCreate(t *testing.T) {
	actual := createConfig(t)
	r, _ := regexp.Compile("[main]")
	if !r.MatchString(actual) {
		t.Errorf("not include %v in %v", r, actual)
	}
	removeTmpFile(t)

}
