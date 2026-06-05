package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff [snapshot-a] [snapshot-b]",
	Short: "Compare two snapshots",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Diff snapshots:", args[0], args[1])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)
}
