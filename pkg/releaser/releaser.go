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
	FileGlob string
	Owner    string
	Repo     string
	TagName  string
	Name     *string
	Body     string
	Draft    bool
}

func (c *ReleaseConfig) String() string {
	return fmt.Sprintf(`FileGlob: %s
Owner: %s
Repo: %s
TagName: %s
Name: %v
Body: %s
Draft: %v`,
		c.FileGlob,
		c.Owner,
		c.Repo,
		c.TagName,
		c.Name,
		c.Body,
		c.Draft,
	)
}

// Release creates a github release
func Release(
	ctx context.Context,
	client *github.Client,
	config ReleaseConfig,
) error {
	fileNames, err := filepath.Glob(config.FileGlob)
	if err != nil {
		return err
	}

	files := make([]*os.File, len(fileNames))
	for i, n := range fileNames {
		file, err := os.OpenFile(n, os.O_RDONLY, 0664)
		if err != nil {
			return err
		}
		files[i] = file
	}

	if config.Name == nil {
		config.Name = &config.TagName
	}

	release, _, err := client.Repositories.CreateRelease(ctx, config.Owner, config.Repo, &github.RepositoryRelease{
		TagName: &config.TagName,
		Name:    config.Name,
		Draft:   &config.Draft,
		Body:    &config.Body,
	})
	if err != nil {
		return err
	}

	for _, f := range files {
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
