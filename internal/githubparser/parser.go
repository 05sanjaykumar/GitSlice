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
	
	if len(parts) < 5 {
		return nil, fmt.Errorf("URL path is too short to parse")
	}

	contentType := parts[3]
	if contentType != "blob" && contentType != "tree" {
		return nil, fmt.Errorf("URL must contain either 'blob' (file) or 'tree' (folder), got: '%s'", contentType)
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