package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]

		err := os.Remove(fileName)
		if err != nil {
			fmt.Println("Error deleting file:", err)
			return
		}

		fmt.Println("File successfully deleted:", fileName)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
