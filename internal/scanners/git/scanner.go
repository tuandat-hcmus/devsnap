package git

import (
	"context"
	"os/exec"
	"strings"
)

type Scanner struct{}

func NewScanner() *Scanner {
	return &Scanner{}
}

func (s *Scanner) Name() string {
	return "git"
}	

func (s *Scanner) Scan(ctx context.Context) (any, error) {
	cmd := exec.CommandContext(ctx, "git", "config", "--global", "--list")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	configs := make(map[string]string)
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			configs[parts[0]] = parts[1]
		}
	}
	return configs, nil
}