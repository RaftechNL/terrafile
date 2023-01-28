package gh

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func DownloadRepository(moduleSource, ref, moduleOutput string) error {

	err := downloadRepositoryByBranch(moduleSource, ref, moduleOutput)
	if err != nil {

		err := downloadRepositoryByTag(moduleSource, ref, moduleOutput)
		if err != nil {
			return err
		}
		return nil

	}
	return nil
}

func downloadRepositoryByTag(moduleSource, ref, moduleOutput string) error {
	os.RemoveAll(moduleOutput)

	refTag := plumbing.NewTagReferenceName(ref)

	_, err := git.PlainClone(moduleOutput, false, &git.CloneOptions{
		URL:           moduleSource,
		SingleBranch:  true,
		ReferenceName: plumbing.ReferenceName(refTag),
	})

	if err != nil {
		return err
	}

	return nil
}

func downloadRepositoryByBranch(moduleSource, ref, moduleOutput string) error {
	os.RemoveAll(moduleOutput)

	refBranch := plumbing.NewBranchReferenceName(ref)

	isBranch := refBranch.IsBranch()
	fmt.Println(isBranch)

	_, err := git.PlainClone(moduleOutput, false, &git.CloneOptions{
		URL:           moduleSource,
		SingleBranch:  true,
		ReferenceName: plumbing.ReferenceName(refBranch),
	})

	if err != nil {
		return err
	}

	return nil
}
