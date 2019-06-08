package travis

import (
	"errors"
	"strings"
)

// RepoSlug contains the owner and repository name of a build
type RepoSlug struct {
	Owner string
	Repo  string
}

// ParseRepoSlug prases a travis repo slug (owner_name/repo_name)
func ParseRepoSlug(str string) (RepoSlug, error) {
	s := strings.SplitN(str, "/", 2)
	if len(s) != 2 {
		return RepoSlug{}, errors.New("could not find any slashes in slug")
	}
	return RepoSlug{
		Owner: s[0],
		Repo:  s[1],
	}, nil
}
