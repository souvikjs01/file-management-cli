package cmd

import (
	"fmt"
	"os"

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

var currentDir string

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting cd", err)
		return
	}
	fmt.Println("current dir", currentDir)
}
