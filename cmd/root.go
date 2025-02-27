package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "file-mgmt-cli",
	Short: "A simple CLI tool for file management",
	Long:  "This tool allows you to perform file operations like list, create, delete, and rename files.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("File Management CLI")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
