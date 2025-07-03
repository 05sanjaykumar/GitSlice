/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/05sanjaykumar/gitslice/internal/githubparser"

	"github.com/05sanjaykumar/gitslice/internal/clone"

	

)

// sliceCmd represents the slice command
var sliceCmd = &cobra.Command{
	Use:   "_default <GitHub URL>",
	Hidden: true,
	Short: "Clone a single file or folder from a GitHub repo using sparse checkout.",
	Long: `Example:
	gitslice https://github.com/user/repo/blob/main/folder/file.js
	`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
        fmt.Println("❌ Please provide a GitHub file or folder URL.")
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

func init() {
	rootCmd.AddCommand(sliceCmd)
}
