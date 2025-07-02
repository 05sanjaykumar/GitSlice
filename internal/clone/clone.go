package slice

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunSparseClone(owner, repo, branch, path string) error {
	repoURL := fmt.Sprintf("https://github.com/%s/%s.git", owner, repo)
	repoName := fmt.Sprintf("%s-%s-temp", repo, filepath.Base(path)) // unique temp folder
	targetName := filepath.Base(path)

	// Step 1: Clone with sparse checkout
	fmt.Println("🚀 Cloning repository...")
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

	// Step 2: Move file/folder out
	src := filepath.Join(repoName, path)
	dst := targetName

	if _, err := os.Stat(src); os.IsNotExist(err) {
		return fmt.Errorf("❌ Target path %s not found in repository", path)
	}

	fmt.Printf("📦 Extracting %s...\n", src)
	err := os.Rename(src, dst)
	if err != nil {
		return fmt.Errorf("❌ Failed to move: %v", err)
	}

	// Step 3: Cleanup
	fmt.Println("🧹 Cleaning up...")
	os.RemoveAll(repoName)

	fmt.Printf("✅ Done! Extracted to: %s\n", dst)
	return nil
}

func runCommand(args ...string) error {
	fmt.Printf("▶️ Running: %s\n", args)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
