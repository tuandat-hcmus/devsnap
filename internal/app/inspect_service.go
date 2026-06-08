package app

import (
	"context"
	"github.com/tuandat-hcmus/devsnap/internal/domain"
)

type InspectService struct {
	Storage domain.SnapshotStorage
}

func NewInspectService(storage domain.SnapshotStorage) *InspectService {
	return &InspectService{
		Storage: storage,
	}
}

func (s *InspectService) GetSnapshot(ctx context.Context, id string) (*domain.Snapshot, error) {
	return s.Storage.FindByID(ctx, id)
}
