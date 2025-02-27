package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rnm",
	Short: "Rename a file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		oldName := args[0]
		newName := args[1]

		err := os.Rename(oldName, newName)
		if err != nil {
			fmt.Println("Error renaming file:", err)
			return
		}

		fmt.Println("File renamed from", oldName, "to", newName)
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
}
