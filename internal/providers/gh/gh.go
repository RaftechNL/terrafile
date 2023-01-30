package gh

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	"github.com/RaftechNL/terrafile/internal/providers"
	"github.com/RaftechNL/terrafile/internal/providers/config"
)

type ProviderGithub struct {
	config *config.ProviderConfig
}

func NewGithubProvider(config *config.ProviderConfig) *ProviderGithub {
	return &ProviderGithub{
		config: config,
	}
}

func (pgh *ProviderGithub) DownloadModule(moduleSpec providers.ModuleSpec, outputPath string) error {

	fmt.Println("Downloading module from Github")
	fmt.Println("Module source: ", moduleSpec.Source)
	fmt.Println("Module ver: ", moduleSpec.Version)

	err := pgh.downloadRepositoryByTag(moduleSpec.Source, moduleSpec.Version, outputPath)
	if err != nil {

		err := pgh.downloadRepositoryByBranch(moduleSpec.Source, moduleSpec.Version, outputPath)
		if err != nil {
			return err
		}
		return nil

	}
	return nil
}

func (pgh *ProviderGithub) downloadRepositoryByTag(moduleSource, ref, moduleOutput string) error {
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

func (pgh *ProviderGithub) downloadRepositoryByBranch(moduleSource, ref, moduleOutput string) error {
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
