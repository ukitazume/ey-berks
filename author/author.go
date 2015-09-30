package author

import (
	"bytes"
	"fmt"
	"github.com/ukitazume/ey-berks/config"
	"os"
	"path/filepath"
)

func getList(berks config.Berks) (list []string) {
	for _, cookbook := range berks.Cookbooks {
		list = append(list, cookbook.RecipeName())
	}
	return
}

func CreateFile(path string, content string) error {
	fullPath := filepath.Join(path, "cookbooks", "main", "recipes")
	fileName := filepath.Join(fullPath, "default.rb")
	if err := os.MkdirAll(fullPath, 0744); err != nil {
		return err
	}
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		return err
	}
	if _, err := f.Write([]byte(content)); err != nil {
		return err
	}
	return nil
}

func CreateMainRecipe(berks config.Berks) string {
	var buffer bytes.Buffer
	descList := "# Created by ey-berks\n"
	buffer.WriteString(descList)
	for _, c := range berks.Cookbooks {
		commentLine := fmt.Sprintf("#from %v\n", c.RemotoRepoUrl())
		buffer.WriteString(commentLine)
		requireLine := fmt.Sprintf("include_recipe \"%s\"\n", c.RecipeName())
		buffer.WriteString(requireLine)
	}
	return buffer.String()
}
