package local

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tuandat-hcmus/devsnap/internal/domain"
	"os"
	"path/filepath"
)

type Storage struct {
	snapshotDir string
}

func NewStorage(snapshotDir string) *Storage {
	return &Storage{snapshotDir: snapshotDir}
}

func (s *Storage) Save(ctx context.Context, snapshot *domain.Snapshot) error {
	if err := os.MkdirAll(s.snapshotDir, 0755); err != nil {
		return fmt.Errorf("failed to create snapshot directory: %w", err)
	}
	data, err := json.MarshalIndent(snapshot, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal snapshot: %w", err)
	}

	filePath := filepath.Join(s.snapshotDir, snapshot.ID+".json")
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write snapshot file: %w", err)
	}
	return nil
}

func (s *Storage) FindByID(ctx context.Context, id string) (*domain.Snapshot, error) {
	filePath := filepath.Join(s.snapshotDir, id+".json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var snapshot domain.Snapshot
	if err := json.Unmarshal(data, &snapshot); err != nil {
		return nil, fmt.Errorf("failed to unmarshal snapshot: %w", err)
	}
	return &snapshot, nil
}

func (s *Storage) List(ctx context.Context) ([]*domain.Snapshot, error) {
	files, err := os.ReadDir(s.snapshotDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read snapshot directory: %w", err)
	}
	var snapshots []*domain.Snapshot
	if len(files) == 0 {
		return snapshots, nil
	}
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}
		id := file.Name()[:len(file.Name())-len(".json")]
		snapshot, err := s.FindByID(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("failed to read snapshot file %s: %w", file.Name(), err)
		}
		snapshots = append(snapshots, snapshot)
	}
	return snapshots, nil
}
