package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"time"

	"github.com/tuandat-hcmus/devsnap/internal/domain"
)

type CaptureService struct {
	SnapshotDir string
	Scanners    []domain.Scanner
}

func NewCaptureService(snapshotDir string, scanners []domain.Scanner) *CaptureService {
	return &CaptureService{
		SnapshotDir: snapshotDir,
		Scanners:    scanners,
	}
}

func (s *CaptureService) saveSnapshot(snapshot *domain.Snapshot) error {
	if err := os.MkdirAll(s.SnapshotDir, 0755); err != nil {
		return fmt.Errorf("failed to create snapshot directory: %w", err)
	}
	data, err := json.MarshalIndent(snapshot, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal snapshot: %w", err)
	}

	filePath := filepath.Join(s.SnapshotDir, snapshot.ID+".json")
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write snapshot file: %w", err)
	}

	return nil
}

func (s *CaptureService) Capture(ctx context.Context, name string) (*domain.Snapshot, error) {
	if name == "" {
		name = "snapshot-" + time.Now().UTC().Format("2006-01-02-15:04:05")
	}
	data := make(map[string]any)
	for _, scanner := range s.Scanners {
		scanData, err := scanner.Scan(context.Background())
		if err != nil {
			data[scanner.Name()] = map[string]any{
				"error": err.Error(),
			}
			continue
		}
		data[scanner.Name()] = scanData
	}
	snapshot := &domain.Snapshot{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
		Data:      data,
	}
	if err := s.saveSnapshot(snapshot); err != nil {
		return nil, err
	}
	return snapshot, nil
}
