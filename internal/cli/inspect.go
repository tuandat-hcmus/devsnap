package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect [snapshot-id]",
	Short: "Inspect a snapshot",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		snapshotID := args[0]

		snapshotPath := filepath.Join("snapshots", snapshotID+".json")
		data, err := os.ReadFile(snapshotPath)
		if err != nil {
			return fmt.Errorf("failed to read snapshot file: %w", err)
		}
		fmt.Println(string(data))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
