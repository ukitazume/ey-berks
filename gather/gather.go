package gather

import (
	"fmt"
	"github.com/libgit2/git2go"
	"github.com/ukitazume/ey-berks/config"
	"os"
	"path/filepath"
)

type Gather struct {
	Berks config.Berks
}

func NewGather(berksFilePath string) Gather {
	bersk := config.Parse(berksFilePath)
	return Gather{Berks: bersk}
}

func (g *Gather) createWorkingDir() error {
	err := os.MkdirAll(workingDir(), 0744)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Join(workingDir(), "github.com"), 0744)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (g *Gather) Gather() error {
	g.createWorkingDir()

	for _, c := range g.Berks.Cookbook {
		prepareDir(c)
		updateCookbook(c)
	}
	return nil
}

func workingDir() string {
	return os.Getenv("HOME") + "/.ey-berks"
}

func prepareDir(cookbook config.Cookbook) error {
	path := filepath.Join(workingDir(), "github.com", cookbook.Repo)
	err := os.MkdirAll(path, 0744)
	if err != nil {
		return err
	}
	return nil
}

func updateCookbook(cookbook config.Cookbook) error {
	path := filepath.Join(workingDir(), "github.com", cookbook.Repo)
	if _, err := os.Stat(path + "/.git"); os.IsNotExist(err) {
		gitCloneOption := new(git.CloneOptions)
		repo, err := git.Clone("git://github.com/"+cookbook.Repo, path, gitCloneOption)
		fmt.Println(repo)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		repo, err := git.OpenRepository(path)
		if err != nil {
			fmt.Println(err)
		}
		remote, err := repo.Remotes.Lookup("origin")
		if err != nil {
			fmt.Println(err)
		}
		err = remote.Fetch([]string{}, nil, "")
		if err != nil {
			fmt.Println(err)
		}

		remoteLs, _ := remote.Ls("HEAD")
		remoteOid := remoteLs[0].Id
		headCommit, _ := repo.LookupCommit(remoteOid)
		fmt.Println(headCommit)
		repo.ResetToCommit(headCommit, git.ResetHard, &git.CheckoutOpts{})
	}
	return nil
}
