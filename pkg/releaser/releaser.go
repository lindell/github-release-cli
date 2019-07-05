package releaser

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/go-github/github"
)

// The ReleaseConfig is the configuration of a release and its uploads
type ReleaseConfig struct {
	FileGlob   string
	Owner      string
	Repo       string
	TagName    string
	Name       string
	Body       string
	Draft      bool
	Prerelease bool

	Logger Logger
}

func (c *ReleaseConfig) String() string {
	return fmt.Sprintf(`FileGlob: %s
Owner: %s
Repo: %s
TagName: %s
Name: %v
Body: %s
Draft: %v
Prerelease: %v`,
		c.FileGlob,
		c.Owner,
		c.Repo,
		c.TagName,
		c.Name,
		c.Body,
		c.Draft,
		c.Prerelease,
	)
}

// Logger contains functions needed for logging
type Logger interface {
	Println(...interface{})
}

func (c *ReleaseConfig) printf(str string, args ...interface{}) {
	if c.Logger == nil {
		return
	}
	c.Logger.Println(fmt.Sprintf(str, args...))
}

// Release creates a github release
func Release(
	ctx context.Context,
	client *github.Client,
	config ReleaseConfig,
) error {
	config.printf("using config:\n%s", config.String())

	fileNames, err := filepath.Glob(config.FileGlob)
	if err != nil {
		return err
	}

	if len(fileNames) == 0 {
		config.printf(`no files found with "%s"`, config.FileGlob)
	} else {
		for _, n := range fileNames {
			config.printf(`found file to upload %s`, n)
		}
	}

	files := make([]*os.File, len(fileNames))
	for i, n := range fileNames {
		file, err := os.OpenFile(n, os.O_RDONLY, 0664)
		if err != nil {
			return err
		}
		files[i] = file
	}

	config.printf("creating a release")
	release, _, err := client.Repositories.CreateRelease(ctx, config.Owner, config.Repo, &github.RepositoryRelease{
		TagName:    &config.TagName,
		Name:       &config.Name,
		Draft:      &config.Draft,
		Body:       &config.Body,
		Prerelease: &config.Prerelease,
	})
	if err != nil {
		return err
	}

	for _, f := range files {
		config.printf("uploading file %s", f.Name())
		_, _, err := client.Repositories.UploadReleaseAsset(ctx, config.Owner, config.Repo, *release.ID, &github.UploadOptions{
			Name:  filepath.Base(f.Name()),
			Label: filepath.Base(f.Name()),
		}, f)
		if err != nil {
			return err
		}
	}

	return nil
}
