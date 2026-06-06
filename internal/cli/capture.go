package cli

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tuandat-hcmus/devsnap/internal/app"
	"github.com/tuandat-hcmus/devsnap/internal/domain"
	gitScanner "github.com/tuandat-hcmus/devsnap/internal/scanners/git"
	systemScanner "github.com/tuandat-hcmus/devsnap/internal/scanners/system"
	vscodeScanner "github.com/tuandat-hcmus/devsnap/internal/scanners/vscode"
)

var snapshotName string
var captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Capture the current development environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		scanners := []domain.Scanner{
			systemScanner.NewScanner(),
			gitScanner.NewScanner(),
			vscodeScanner.NewScanner(),
		}
		service := app.NewCaptureService("snapshots", scanners)
		snapshot, err := service.Capture(context.Background(), snapshotName)
		if err != nil {
			fmt.Println("Error capturing snapshot:", err)
			return err
		}

		fmt.Printf("Snapshot captured: %s\n", snapshot.ID)
		fmt.Printf("Name: %s\n", snapshot.Name)
		return nil
	},
}

func init() {
	captureCmd.Flags().StringVarP(&snapshotName, "name", "n", "", "Name of the snapshot")
	rootCmd.AddCommand(captureCmd)
}
