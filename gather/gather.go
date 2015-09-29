package gather

import (
	"fmt"
	"github.com/libgit2/git2go"
	"github.com/ukitazume/ey-berks/config"
	"os"
	"os/exec"
	"path/filepath"
)

type Gather struct {
	Berks config.Berks
}

func NewGather(berksFilePath string) Gather {
	bersk := config.Parse(berksFilePath)
	return Gather{Berks: bersk}
}

func (g *Gather) Gather() error {
	prepareCookBookDir("./")

	for _, c := range g.Berks.Cookbook {
		prepareDir(c)
		if err := updateCookbook(c); err != nil {
			fmt.Println(err)
		}
		copyRecipes("./cookbooks", c)
	}
	return nil
}

func prepareDir(cookbook config.Cookbook) error {
	err := os.MkdirAll(cookbook.WorkingRepoPath(), 0744)
	if err != nil {
		return err
	}
	return nil
}

func updateCookbook(cookbook config.Cookbook) error {
	path := cookbook.WorkingRepoPath()
	if _, err := os.Stat(path + "/.git"); os.IsNotExist(err) {
		gitCloneOption := new(git.CloneOptions)
		if _, err := git.Clone(cookbook.RemoteUrl(), path, gitCloneOption); err != nil {
			return err
		}
		fmt.Printf("clone to locat from %s\n", cookbook.RemoteUrl())
	} else {
		repo, err := git.OpenRepository(path)
		if err != nil {
			return err
		}

		remote, err := repo.Remotes.Lookup("origin")
		fmt.Printf("fetching to locat from %s\n", remote.Url())
		if err != nil {
			return err
		}
		if err := remote.Fetch([]string{}, nil, ""); err != nil {
			return err
		}

		remoteLs, err := remote.Ls("HEAD")
		if err != nil {
			return err
		}
		remoteOid := remoteLs[0].Id
		headCommit, err := repo.LookupCommit(remoteOid)
		if err != nil {
			return err
		}
		if err := repo.ResetToCommit(headCommit, git.ResetHard, &git.CheckoutOpts{}); err != nil {
			return err
		}
		fmt.Printf(" -- now %s\n", remoteOid)
	}
	return nil
}

func prepareCookBookDir(path string) error {
	fullPath := filepath.Join(path, "cookbooks")
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		if err := os.MkdirAll(fullPath, 0744); err != nil {
			return err
		}
	}
	return nil
}

func copyRecipes(path string, c config.Cookbook) error {
	destDir := filepath.Join(path, "cookbooks", c.Host(), c.Repo, c.Path)
	srcDir := c.WorkingPath()
	exec.Command("cp", "-rf", srcDir, destDir)
	return nil
}
