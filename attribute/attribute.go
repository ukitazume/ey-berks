package attribute

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Options struct {
	TempDirName string
}

const (
	TempDirName = "attr-meta"
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
	fmt.Println("searching into %v", searchPath)
	attributes, err := filepath.Glob(searchPath)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(pathDist, opts.TempDirName), 0755); err != nil {
		return err
	}
	for _, attr := range attributes {
		dirs := strings.Split(attr, "/")
		newDirName := filepath.Join(opts.distPath(pathDist), dirs[len(dirs)-2])
		fmt.Printf("find %v\n", newDirName)
		if err := os.MkdirAll(opts.distPath(pathDist), 0755); err != nil {
			return err
		}
		if err := copyFile(attr, newDirName); err != nil {
			return err
		}
	}
	return nil
}

func Apply(path string, pathAttributes string, opts Options) error {
	attributes, err := filepath.Glob(filepath.Join(pathAttributes, "**", "*.rb"))
	if err != nil {
		return err
	}
	for _, attr := range attributes {
		dirs := strings.Split(attr, "/")
		distRecipeDir := filepath.Join(path, "cookbooks", dirs[len(dirs)-2])
		distDir := filepath.Join(distRecipeDir, "attributes")
		dist := filepath.Join(distDir, dirs[len(dirs)-1])
		if _, err := os.Stat(distRecipeDir); os.IsNotExist(err) {
			fmt.Printf("- skip %v because no recipes in cookbooks\n", attr)
			continue
		}
		if err := os.MkdirAll(distDir, 0755); err != nil {
			return err
		}
		fmt.Printf("+ copy to %v to %v\n", attr, dist)
		if err := copyFile(attr, dist); err != nil {
			return err
		}
	}
	return nil
}

func copyFile(srcPath, distPath string) error {
	cpCmd := exec.Command("cp", "-rf", srcPath, distPath)
	if err := cpCmd.Run(); err != nil {
		return err
	}
	return nil
}
