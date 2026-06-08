package domain

import "context"

type SnapshotStorage interface {
	Save(ctx context.Context, snapshot *Snapshot) error
	FindByID(ctx context.Context, id string) (*Snapshot, error)
	List(ctx context.Context) ([]*Snapshot, error)
}
