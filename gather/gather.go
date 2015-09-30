package gather

import (
	"fmt"
	"github.com/libgit2/git2go"
	"github.com/ukitazume/ey-berks/config"
	"os"
	"os/exec"
	"path/filepath"
)

func Gather(path string, berks config.Berks) error {
	if err := gatherDir(berks.Library, path); err != nil {
		fmt.Printf("error: %v", err)
	}
	if err := gatherDir(berks.Definition, path); err != nil {
		fmt.Printf("error: %v", err)
	}

	for _, cookbook := range berks.Cookbooks {
		if err := gatherDir(cookbook, path); err != nil {
			fmt.Printf("error: %v", err)
		}
	}

	return nil
}

func Copy(path string, berks config.Berks) error {
	if err := copyRecipes(berks.Library, path); err != nil {
		fmt.Printf("error: %v", err)
	}
	if err := copyRecipes(berks.Definition, path); err != nil {
		fmt.Printf("error: %v", err)
	}

	for _, cookbook := range berks.Cookbooks {
		if err := copyRecipes(cookbook, path); err != nil {
			fmt.Printf("error: %v", err)
		}
	}

	return nil
}

func gatherDir(c config.CodeResourceOperator, path string) error {
	if err := prepareDir(c, path); err != nil {
		return err
	}
	if err := updateCookbook(c); err != nil {
		return err
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
		fmt.Println(c.RemotoRepoUrl())
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
		// TODO use short ref id
		fmt.Printf(" -- now %s\n", remoteOid)
	}
	return nil
}

func copyRecipes(c config.CodeResourceOperator, path string) error {
	if err := prepareDir(c, path); err != nil {
		return err
	}
	destDir := filepath.Join(path, c.DesticationPath())
	srcDir := filepath.Join(c.CacheRepoPath(), c.SourcePath())
	fmt.Printf(" -- copy %s to %s\n", srcDir, destDir)
	cmd := exec.Command("cp", "-rf", srcDir+"/", destDir)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
