package githubparser

import (
	"fmt"
	"net/url"
	"strings"
)

type GitHubURL struct {
	Owner    string
	Repo     string
	PostTree []string
}

func Parse(rawURL string) (*GitHubURL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	parts := strings.Split(u.Path, "/")
	if len(parts) < 6 {
		return nil, fmt.Errorf("URL path too short to parse")
	}

	if parts[3] != "tree" && parts[3] != "blob" {
		return nil, fmt.Errorf("URL must contain 'tree' or 'blob'")
	}

	return &GitHubURL{
		Owner:    parts[1],
		Repo:     parts[2],
		PostTree: parts[4:],
	}, nil
}
