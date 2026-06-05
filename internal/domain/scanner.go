package domain

import "context"

type Scanner interface {
	Name() string
	Scan(ctx context.Context) (any, error)
}