package gather

import (
	"github.com/ukitazume/ey-berks/config"
	"os"
	"path/filepath"
	"testing"
)

func TestNewGather(t *testing.T) {
	if err := config.Create("/tmp"); err != nil {
		t.Errorf("cannot create Berksfile with %d", err)
	}
	gather := NewGather("/tmp")
	err := gather.Gather("/tmp")
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
