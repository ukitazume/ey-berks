package attribute

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Options struct {
	TempDirName string
}

const (
	TempDirName = "attributes"
)

func DefaultOptions() Options {
	return Options{TempDirName}
}

func (opts Options) searchPath(path string) string {
	return filepath.Join(path, "**", "attributes")
}

func (opts Options) distPath(path string) string {
	return filepath.Join(path, opts.TempDirName)
}

func Gather(pathOrignal string, pathDist string, opts Options) error {
	searchPath := opts.searchPath(pathOrignal)
	attributes, err := filepath.Glob(searchPath)
	if err != nil {
		return err
	}
	os.MkdirAll(filepath.Join("./", opts.TempDirName), 0755)
	for _, attr := range attributes {
		dirs := strings.Split(attr, "/")
		newDirName := filepath.Join(opts.distPath(pathDist), dirs[len(dirs)-2])
		if err := os.MkdirAll(opts.distPath(pathDist), 0755); err != nil {
			return err
		}
		if err := copyFile(attr, newDirName); err != nil {
			return err
		}
	}
	return nil
}

func Apply(path string, opt Options) error {
	return nil
}

func copyFile(srcPath, distPath string) error {
	cpCmd := exec.Command("cp", "-rf", srcPath, distPath)
	if err := cpCmd.Run(); err != nil {
		return err
	}
	return nil
}
