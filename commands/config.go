package config

import (
	"bytes"
	"os"
)

func Create(path string) error {
	var fullPath bytes.Buffer
	fullPath.WriteString(path)
	fullPath.WriteString("/EYBerksfile")

	f, err := os.Create(fullPath.String())
	if err != nil {
		return err
	}
	defaultFormat := `[main]
libraries = "engineyard/ey-cloud-recipes/main/libraries"
definitions = "engineyard/ey-cloud-recipes/main/definitions"

[cookbook]
env_vars = "engineyard/ey-cloud-recipes/cookbooks/env_vars"
`
	f.Write([]byte(defaultFormat))
	f.Close()

	return nil
}
