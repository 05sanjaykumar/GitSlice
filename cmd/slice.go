/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/05sanjaykumar/Gitslice-CLI/internal/githubparser"

	"github.com/05sanjaykumar/Gitslice-CLI/internal/clone"

	

)

// sliceCmd represents the slice command
var sliceCmd = &cobra.Command{
	Use:   "gitslice <GitHub URL>",
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sliceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sliceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
