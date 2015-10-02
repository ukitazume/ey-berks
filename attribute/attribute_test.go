package attribute

import (
	"os"
	"path/filepath"
	"testing"
)

func prepare(t *testing.T) string {
	fullPath := filepath.Join(os.TempDir(), "ey-berks", "cookbooks")
	attributesPath := filepath.Join(fullPath, "nginx", "attributes")
	if err := os.MkdirAll(attributesPath, 0744); err != nil {
		t.Fatal("cannot create test directory")
	}
	attrs := []string{"default.rb", "nginx.rb"}
	for _, v := range attrs {
		filePath := filepath.Join(attributesPath, v)
		if _, err := os.Create(filePath); err != nil {
			t.Fatalf("error %v", err)
		}
	}
	return fullPath
}

func tearDown() {
	os.RemoveAll(filepath.Join(os.TempDir(), "ey-berks"))
}

func TestGather(t *testing.T) {
	pathOriginal := prepare(t)
	pathTarget := filepath.Join(os.TempDir(), "ey-berks")
	opts := DefaultOptions()
	if err := Gather(pathOriginal, pathTarget, opts); err != nil {
		t.Fatalf("error %v", err)
	}
	filename := filepath.Join(os.TempDir(), "ey-berks", opts.TempDirName, "nginx", "default.rb")
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("error: file doesn't exist  %v", err)
	}

	tearDown()
}
