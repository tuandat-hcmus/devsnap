package app

import (
	"context"
	"github.com/tuandat-hcmus/devsnap/internal/domain"
)

type ListService struct {
	Storage domain.SnapshotStorage
}

func NewListService(storage domain.SnapshotStorage) *ListService {
	return &ListService{
		Storage: storage,
	}
}

func (s *ListService) List(ctx context.Context) ([]*domain.Snapshot, error) {
	return s.Storage.List(ctx)
}
