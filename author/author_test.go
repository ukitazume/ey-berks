package author

import (
	"fmt"
	"github.com/ukitazume/ey-berks/config"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func prepare(t *testing.T) config.Berks {
	if err := config.Create("/tmp"); err != nil {
		t.Errorf("cannot create Berksfile with %d", err)
	}
	berks := config.Parse("/tmp")
	return berks
}

func tearDown(t *testing.T) {
	path := filepath.Join("/tmp", config.ConfigFileName)
	if err := os.Remove(path); err != nil {
		t.Errorf("failed to remove %s", path)
	}
}

func TestCreateFile(t *testing.T) {
	berks := prepare(t)
	list := createMainRecipe(berks)
	if err := createFile("/tmp", list); err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestCreateMainRecipe(t *testing.T) {
	berks := prepare(t)
	list := createMainRecipe(berks)

	for _, v := range []string{"env_vars", "custom_nginx"} {
		requireLine := fmt.Sprintf("include_recipe \"%s\"", v)
		r, _ := regexp.Compile(requireLine)
		if !r.MatchString(list) {
			t.Errorf("not include %v in %v", r, list)
		}
	}
	tearDown(t)
}
