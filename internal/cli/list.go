package cli

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tuandat-hcmus/devsnap/internal/app"
	"github.com/tuandat-hcmus/devsnap/internal/storage/local"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snapshots",
	RunE: func(cmd *cobra.Command, args []string) error {
		storage := local.NewStorage("snapshots")
		listService := app.NewListService(storage)
		ctx := context.Background()
		snapshots, err := listService.List(ctx)
		if err != nil {
			fmt.Println("Error listing snapshots:", err)
			return err
		}
		for _, snapshot := range snapshots {
			fmt.Println(snapshot.ID)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
