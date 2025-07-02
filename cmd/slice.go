/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sliceCmd represents the slice command
var sliceCmd = &cobra.Command{
	Use:   "slice",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
        fmt.Println("❌ Please provide a GitHub file or folder URL.")
        return
    }

		gh, err := parseGitHubURL(args[0])
		if err != nil {
			fmt.Printf("❌ Error parsing URL: %v\n", err)
			return
		}

		err = slice.RunSparseClone(gh.Owner, gh.Repo, gh.Branch, gh.Path)
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
