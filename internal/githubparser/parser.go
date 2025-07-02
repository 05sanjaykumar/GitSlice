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

	parts := strings.Split(strings.Trim(u.Path, "/"), "/") // Trim leading/trailing slashes
	if len(parts) < 2 {
		return nil, fmt.Errorf("URL path too short to parse")
	}

	// Full repo clone URL: https://github.com/owner/repo
	if len(parts) == 2 {
		return &GitHubURL{
			Owner:    parts[0],
			Repo:     strings.TrimSuffix(parts[1], ".git"),
			PostTree: []string{}, // no subdir
		}, nil
	}

	// Sparse clone URL: https://github.com/owner/repo/tree/branch/path/to/folder
	if len(parts) >= 4 && (parts[2] == "tree" || parts[2] == "blob") {
		return &GitHubURL{
			Owner:    parts[0],
			Repo:     parts[1],
			PostTree: parts[3:], // includes branch + path
		}, nil
	}

	return nil, fmt.Errorf("URL must be a GitHub repo or contain 'tree' or 'blob' for paths")
}
