package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "document",
	Short: "Document Service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from myapp!")
	},
}

func Execute() {
	rootCmd.AddCommand(serveCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
