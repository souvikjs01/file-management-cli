package main

import (
	"file-mgmt-cli/cmd"
	"fmt"
	"os"
)

func main() {

	if err := cmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
