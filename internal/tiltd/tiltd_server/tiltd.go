package tiltd_server

import (
	"context"
	"os"
	"os/exec"
)

func RunDaemon(ctx context.Context) (*os.Process, error) {
	cmd := exec.CommandContext(ctx, os.Args[0], "daemon")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	return cmd.Process, nil
}
