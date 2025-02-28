package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [directory]",
	Short: "List files in a directory",
	Run: func(cmd *cobra.Command, args []string) {
		dir := currentDir

		if len(args) > 0 {
			dir = args[0]
		}

		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println("error reading the directory:", err)
			return
		}

		if len(files) == 0 {
			fmt.Println("No files found.")
			return
		}
		fmt.Println("files in", dir, ":")
		for _, file := range files {
			fmt.Println(file.Name())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
