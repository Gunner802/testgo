package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "testcli",
		Short: "This is a test CLI",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Cobra!")
		},
	}

	rootCmd.Execute()
}
