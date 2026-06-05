package system

import (
	"context"
	"os"
	"runtime"
)

type Scanner struct{}

type SystemInfo struct {
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	Hostname string `json:"hostname"`
	NumCPU   int    `json:"num_cpu"`
}

func NewScanner() *Scanner {
	return &Scanner{}
}

func (s *Scanner) Name() string {
	return "system"
}

func (s *Scanner) Scan(ctx context.Context) (any, error) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	info := &SystemInfo{
		OS:       runtime.GOOS,
		Arch:     runtime.GOARCH,
		Hostname: hostname,
		NumCPU:   runtime.NumCPU(),
	}
	return info, nil
}
