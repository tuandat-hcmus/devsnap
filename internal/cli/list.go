package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snapshots",
	RunE: func(cmd *cobra.Command, args []string) error {
		files, err := os.ReadDir("snapshots")
		if err != nil {
			return fmt.Errorf("failed to read snapshots directory: %w", err)
		}
		if len(files) == 0 {
			fmt.Println("No snapshots found.")
			return nil
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if filepath.Ext(file.Name()) != ".json" {
				continue
			}
			fmt.Println(file.Name())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
