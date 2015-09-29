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

func (g *Gather) Gather(path string) error {
	prepareDir(g.Berks.Library, path)
	updateCookbook(g.Berks.Library)
	copyRecipes(g.Berks.Library, path)

	prepareDir(g.Berks.Definition, path)
	updateCookbook(g.Berks.Definition)
	for _, cookbook := range g.Berks.Cookbooks {
		prepareDir(cookbook, path)
	}

	return nil
}

func prepareDir(c config.CodeResourceOperator, path string) error {
	fullPath := filepath.Join(path, c.DesticationPath())
	fmt.Printf("create %v\n", fullPath)

	err := os.MkdirAll(fullPath, 0744)
	if err != nil {
		return err
	}
	return nil
}

func updateCookbook(c config.CodeResourceOperator) error {
	path := c.CacheRepoPath()
	if _, err := os.Stat(path + "/.git"); os.IsNotExist(err) {
		gitCloneOption := new(git.CloneOptions)
		if _, err := git.Clone(c.RemotoRepoUrl(), path, gitCloneOption); err != nil {
			return err
		}
		fmt.Printf("clone to locat from %s\n", c.RemotoRepoUrl())
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

func copyRecipes(c config.CodeResourceOperator, path string) error {
	destDir := filepath.Join(path, c.DesticationPath())
	srcDir := filepath.Join(c.CacheRepoPath(), c.SourcePath())
	fmt.Printf(" -- copy %s to %s\n", srcDir, destDir)
	cmd := exec.Command("cp", "-rf", srcDir+"/", destDir)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
