package github

import (
	_ "embed"
	"context"
	// "io/ioutil"
	"path/filepath"

	"github.com/Jon1105/pmag/utilities"
	"github.com/Jon1105/pmag/vcs/git"

	"github.com/google/go-github/v34/github"
	"golang.org/x/oauth2"
)

//go:embed github.key
var key string

func createRepo(repoName string, private bool) (*github.Repository, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: key},
	)

	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)
	repo := &github.Repository{
		Name:    github.String(repoName),
		Private: github.Bool(private),
	}
	resRepo, _, err2 := client.Repositories.Create(context.Background(), "", repo)
	return resRepo, err2
}

func Github(projectPath string, private bool) error {
	if err1 := git.Git(projectPath); err1 != nil {
		return err1
	}
	repo, err2 := createRepo(filepath.Base(projectPath), private)
	if err2 != nil {
		return err2
	}
	return utilities.RunCommand(projectPath, "git", "remote", "add", "origin", *repo.CloneURL)
}
