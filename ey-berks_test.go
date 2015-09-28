package main

import (
	"testing"
)

func TestNoArgs(t *testing.T) {
	result := Command(nil)
	expect := 0
	if result != expect {
		t.Errorf("got %v\nwant %v", result, expect)
	}
}
