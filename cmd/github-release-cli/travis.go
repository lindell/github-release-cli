package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/lindell/github-release-cli/pkg/releaser"
	"github.com/lindell/github-release-cli/pkg/travis"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_OATH_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repoSlug, err := travis.ParseRepoSlug(os.Getenv("TRAVIS_REPO_SLUG"))
	if err != nil {
		log.Fatal(err)
	}

	tag := os.Getenv("TRAVIS_TAG")

	var name string
	if tag != "" {
		name = tag
	} else {
		name = fmt.Sprintf("%s-%s", os.Getenv("TRAVIS_BRANCH"), os.Getenv("TRAVIS_COMMIT"))
	}

	config := releaser.ReleaseConfig{
		FileGlob: os.Getenv("FILES"),
		Owner:    repoSlug.Owner,
		Repo:     repoSlug.Repo,
		TagName:  name,
		Body:     os.Getenv("BODY"),
	}

	fmt.Println(config.String())

	err = releaser.Release(
		ctx,
		client,
		config,
	)
	if err != nil {
		log.Fatal(err)
	}
}
