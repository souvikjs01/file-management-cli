package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirName := args[0]
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			fmt.Println("error creating directory:", err)
		}
		fmt.Println("Directory created:", dirName)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
