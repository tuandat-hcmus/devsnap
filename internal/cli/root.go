package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "devsnap",
	Short: "DevSnap is a CLI tool to capture and restore development environment snapshots",
	Long:  `DevSnap allows developers to capture the state of their development environment, including installed tools, configurations, and project dependencies.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Welcome to DevSnap!")
		fmt.Println("Run 'devsnap --help' to see available commands.")
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
