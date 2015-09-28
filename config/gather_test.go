package gather

import (
	"testing"
)

func TestNewGather(t *testing.T) {
	gather := NewGather("/tmp/EYBerksfile")
	gather.Gather()
}
