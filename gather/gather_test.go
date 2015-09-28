package gather

import (
	"github.com/ukitazume/ey-berks/config"
	"testing"
)

func TestNewGather(t *testing.T) {
	if err := config.Create("/tmp"); err != nil {
		t.Errorf("cannot create Berksfile with %d", err)
	}
	gather := NewGather("/tmp/EYBerksfile")
	err := gather.Gather()
	if err != nil {
		t.Errorf("%d", err)
	}
}
