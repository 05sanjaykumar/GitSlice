package githubparser

import (
	"fmt"
	"net/url"
	"strings"
)

type GitHubURL struct {
	Owner  string
	Repo   string
	Branch string
	Path   string
}


func Parse(rawURL string) (*GitHubURL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	parts := strings.Split(u.Path, "/")
	if len(parts) < 5 || parts[3] != "blob" {
		return nil, fmt.Errorf("URL does not match expected GitHub blob pattern")
	}

	owner := parts[1]
	repo := parts[2]
	branch := parts[4]
	path := strings.Join(parts[5:], "/")

	return &GitHubURL{
		Owner:  owner,
		Repo:   repo,
		Branch: branch,
		Path:   path,
	}, nil
}