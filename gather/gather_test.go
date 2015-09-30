package gather

import (
	"github.com/ukitazume/ey-berks/config"
	"os"
	"path/filepath"
	"testing"
)

func TestNewGather(t *testing.T) {
	opts := config.DefaultOption()
	if err := config.Create("/tmp", opts); err != nil {
		t.Errorf("cannot create Berksfile with %d", err)
	}
	berks := config.Parse("/tmp", opts)
	err := Gather("/tmp", berks)
	if err != nil {
		t.Errorf("%d", err)
	}

	removeTmpFile(t)
}

func removeTmpFile(t *testing.T) {
	path := filepath.Join("/tmp", config.ConfigFileName)
	if err := os.Remove(path); err != nil {
		t.Errorf("failed to remove %s", path)
	}
	if err := os.RemoveAll("/tmp/cookbooks"); err != nil {
		t.Errorf("failed to remove /tmp/cookbooks")
	}
}
