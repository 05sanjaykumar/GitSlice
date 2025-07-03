/*
Copyright © 2025 Sanjay Kumar S

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/05sanjaykumar/gitslice/internal/clone"

	"github.com/05sanjaykumar/gitslice/internal/githubparser"
	"github.com/spf13/cobra"
)

var version = "dev" // overridden by goreleaser

var rootCmd = &cobra.Command{
	Use:     "gitslice <GitHub URL>",
	Short:   "⚡ Extract folders from GitHub repos — fast",
	Long: `GitSlice is a blazing-fast CLI tool to clone or extract any folder or file from a GitHub repository using sparse-checkout.

Examples:
  # Extract a subfolder/file from a GitHub repo
  gitslice https://github.com/user/repo/tree/main/path/to/<folder_or_file>

  # Perform a full clone (no path specified)
  gitslice https://github.com/user/repo

Features:
  • Sparse or full clone support
  • Auto path resolution
  • Cross-platform compatible`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}

		gh, err := githubparser.Parse(args[0])
		if err != nil {
			fmt.Printf("❌ Error parsing URL: %v\n", err)
			return
		}

		err = slice.RunSparseClone(gh.Owner, gh.Repo, gh.PostTree)
		if err != nil {
			fmt.Printf("❌ Clone failed: %v\n", err)
			return
		}

		fmt.Println("✅ Done.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
}