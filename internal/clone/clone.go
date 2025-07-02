package slice

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunSparseClone(owner, repo string, postTree []string) error {
	repoURL := fmt.Sprintf("https://github.com/%s/%s.git", owner, repo)
	cloneTemp := fmt.Sprintf("%s-branch-resolve-temp", repo)

	// Step 0: Shallow clone without checkout to detect valid branch and path
	fmt.Println("🚀 Cloning repository for branch resolution...")
	cmd := exec.Command("git", "clone", "--filter=blob:none", "--no-checkout", repoURL, cloneTemp)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("❌ Failed to clone for resolving branch: %w", err)
	}

	branch, path, err := resolveBranchAndPath(cloneTemp, postTree)
	if err != nil {
		os.RemoveAll(cloneTemp)
		return err
	}
	os.RemoveAll(cloneTemp) // Clean temp clone

	// Unique temp folder for sparse checkout
	repoName := fmt.Sprintf("%s-%s-temp", repo, filepath.Base(path))
	targetName := filepath.Base(path)

	// Step 1: Sparse checkout
	commands := [][]string{
		{"git", "clone", "--filter=blob:none", "--no-checkout", repoURL, repoName},
		{"git", "-C", repoName, "sparse-checkout", "init", "--no-cone"},
		{"git", "-C", repoName, "sparse-checkout", "set", path},
		{"git", "-C", repoName, "checkout", branch},
	}

	for _, cmd := range commands {
		if err := runCommand(cmd...); err != nil {
			return err
		}
	}

	// Step 2: Move file/folder
	src := filepath.Join(repoName, path)
	dst := targetName

	if _, err := os.Stat(src); os.IsNotExist(err) {
		return fmt.Errorf("❌ Target path %s not found in repository", path)
	}

	fmt.Printf("📦 Extracting %s...\n", src)
	if err := os.Rename(src, dst); err != nil {
		return fmt.Errorf("❌ Failed to move: %v", err)
	}

	// Step 3: Cleanup
	fmt.Println("🧹 Cleaning up...")
	os.RemoveAll(repoName)

	fmt.Printf("✅ Done! Extracted to: %s\n", dst)
	return nil
}


func resolveBranchAndPath(clonePath string, postTree []string) (string, string, error) {
	for i := 1; i < len(postTree); i++ {
		branchCandidate := postTree[:i]
		pathCandidate := postTree[i:]

		if len(pathCandidate) == 0 {
			continue
		}

		// Join path inside clone directory
		fullPath := filepath.Join(clonePath, filepath.Join(pathCandidate...))

		if info, err := os.Stat(fullPath); err == nil && info.IsDir() {
			// We found a valid path inside repo, return match
			return strings.Join(branchCandidate, "/"), strings.Join(pathCandidate, "/"), nil
		}
	}

	return "", "", fmt.Errorf("❌ Could not resolve branch/path. Check if URL or folders exist")
}

func runCommand(args ...string) error {
	fmt.Printf("▶️ Running: %s\n", args)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
