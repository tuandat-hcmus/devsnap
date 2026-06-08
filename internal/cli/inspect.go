package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tuandat-hcmus/devsnap/internal/app"
	"github.com/tuandat-hcmus/devsnap/internal/storage/local"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect [snapshot-id]",
	Short: "Inspect a snapshot",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		snapshotID := args[0]
		storage := local.NewStorage("snapshots")
		inspectService := app.NewInspectService(storage)
		ctx := context.Background()
		snapshot, err := inspectService.GetSnapshot(ctx, snapshotID)
		if err != nil {
			fmt.Println("Error inpecting snapshot")
			return err
		}
		jsonData, err := json.MarshalIndent(snapshot, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling snapshot to JSON")
			return err
		}
		fmt.Println(string(jsonData))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
