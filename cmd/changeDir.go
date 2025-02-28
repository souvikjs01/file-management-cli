package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cdCmd = &cobra.Command{
	Use:   "cdr [directory]",
	Short: "Change current directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newDir := args[0]

		err := os.Chdir(newDir)
		if err != nil {
			fmt.Println("error changing the directory")
			return
		}

		// update global current directory:
		currentDir, _ = os.Getwd()
		fmt.Println("current directory changed to:", currentDir)
	},
}

func init() {
	rootCmd.AddCommand(cdCmd)
}
