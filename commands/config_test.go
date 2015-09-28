package config

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

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
	os.Remove("/tmp/EYBerksfile")
}
