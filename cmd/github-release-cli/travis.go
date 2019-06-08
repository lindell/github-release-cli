package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/lindell/github-release-cli/pkg/releaser"
	"github.com/lindell/github-release-cli/pkg/travis"
	"golang.org/x/oauth2"
)

var (
	draft   = false
	verbose = false
)

func main() {
	flag.BoolVar(&draft, "draft", false, "set if the the release should be added as a draft")
	flag.BoolVar(&verbose, "verbose", false, "print logging statements")
	flag.Parse()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_OAUTH_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repoSlug, err := travis.ParseRepoSlug(os.Getenv("TRAVIS_REPO_SLUG"))
	if err != nil {
		log.Fatal(err)
	}

	tagEnv := os.Getenv("TRAVIS_TAG")
	nameEnv := os.Getenv("RELEASE_NAME")

	var tag string
	var name string

	if tagEnv != "" {
		tag = tagEnv
	} else {
		tag = fmt.Sprintf("%s-%s", os.Getenv("TRAVIS_BRANCH"), os.Getenv("TRAVIS_COMMIT"))
	}

	if nameEnv != "" {
		name = nameEnv
	} else {
		name = tag
	}

	var logger releaser.Logger
	if verbose {
		logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}

	config := releaser.ReleaseConfig{
		FileGlob: os.Getenv("FILES"),
		Owner:    repoSlug.Owner,
		Repo:     repoSlug.Repo,
		TagName:  tag,
		Name:     name,
		Body:     os.Getenv("BODY"),
		Draft:    draft,
		Logger:   logger,
	}

	err = releaser.Release(
		ctx,
		client,
		config,
	)
	if err != nil {
		log.Fatal(err)
	}
}
