package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore [snapshot-id]",
	Short: "Restore development environment from snapshot",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Restoring snapshot:", args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}
