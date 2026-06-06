package vscode

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
	return "vscode"
}

func (s *Scanner) Scan(ctx context.Context) (any, error) {
	cmd := exec.CommandContext(ctx, "code", "--list-extensions")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	data := make(map[string]any)
	lines := strings.Split(string(output), "\n")
	data["extensions"] = lines[:len(lines)-1] // Remove the last empty line
	return data, nil
}
