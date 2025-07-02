package slice

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunSparseClone(owner, repo string, postTree []string) error {
	fmt.Printf("owner: %s, repo: %s, postTree: %s\n", owner, repo, strings.Join(postTree, "/"))
	// owner: supabase, repo: storage, postTree: fix/pgboss-on-error-callback/src/auth
	repoURL := fmt.Sprintf("https://github.com/%s/%s.git", owner, repo)
	// https://github.com/supabase/storage.git
	cloneTemp := fmt.Sprintf("%s-branch-resolve-temp", repo)
	// storage-branch-resolve-temp

	// Step 0: Shallow clone without checkout to detect valid branch and path
	fmt.Println("🚀 Cloning repository for branch resolution...")
	cmd := exec.Command("git", "clone", "--depth", "1", repoURL, cloneTemp)
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


// cloneTemp: storage-branch-resolve-temp
// postTree: [fix, pgboss-on-error-callback, src, auth
func resolveBranchAndPath(clonePath string, postTree []string) (string, string, error) {
	// Step 1: Read top-level folders inside the cloned repo
	files, err := os.ReadDir(clonePath)

	fmt.Println("📁 Top-level folders in cloned repo:")
	for _, f := range files {
		if f.IsDir() {
			fmt.Println("  -", f.Name())
		}
	}

	if err != nil {
		return "", "", fmt.Errorf("❌ Failed reading cloned folder: %v", err)
	}

	// Step 2: Store folder names in a set (for O(1) lookup)
	folderSet := make(map[string]bool)
	for _, file := range files {
		if file.IsDir() {
			folderSet[file.Name()] = true
		}
	}

	// Step 3: Find the index in postTree where a match begins
	for i := 0; i < len(postTree); i++ {
		if folderSet[postTree[i]] {
			// Found the beginning of a valid path
			branchCandidate := strings.Join(postTree[:i], "/")
			pathCandidate := strings.Join(postTree[i:], "/")

			// Confirm that the full path actually exists
			fullPath := filepath.Join(clonePath, filepath.Join(postTree[i:]...))
			if info, err := os.Stat(fullPath); err == nil && info.IsDir() {
				return branchCandidate, pathCandidate, nil
			}
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
