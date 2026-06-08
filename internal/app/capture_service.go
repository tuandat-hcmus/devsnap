package app

import (
	"context"
	"github.com/google/uuid"
	"sync"
	"time"

	"github.com/tuandat-hcmus/devsnap/internal/domain"
)

type CaptureService struct {
	Storage  domain.SnapshotStorage
	Scanners []domain.Scanner
}

type ScanResult struct {
	Name string
	Data any
	Err  error
}

func NewCaptureService(storage domain.SnapshotStorage, scanners []domain.Scanner) *CaptureService {
	return &CaptureService{
		Storage:  storage,
		Scanners: scanners,
	}
}

func (s *CaptureService) Capture(ctx context.Context, name string) (*domain.Snapshot, error) {
	if name == "" {
		name = "snapshot-" + time.Now().UTC().Format("2006-01-02-15:04:05")
	}
	data := s.runScanners(ctx)
	snapshot := &domain.Snapshot{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
		Data:      data,
	}
	if err := s.Storage.Save(ctx, snapshot); err != nil {
		return nil, err
	}
	return snapshot, nil
}

func (s *CaptureService) runScanners(ctx context.Context) map[string]any {
	results := make(map[string]any)
	resultCh := make(chan ScanResult)
	var wg sync.WaitGroup
	for _, scanner := range s.Scanners {
		wg.Add(1)
		go func(scanner domain.Scanner) {
			defer wg.Done()
			data, err := scanner.Scan(ctx)
			resultCh <- ScanResult{
				Name: scanner.Name(),
				Data: data,
				Err:  err,
			}
		}(scanner)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh {
		if result.Err != nil {
			results[result.Name] = map[string]any{
				"error": result.Err.Error(),
			}
			continue
		}
		results[result.Name] = result.Data
	}
	return results
}
