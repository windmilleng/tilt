package k8s

import (
	"bytes"
	"context"
	"os"
	"os/exec"
)

func Apply(ctx context.Context, rawYAML string) error {
	// TODO(dmiller) validate that the string is YAML and give a good error
	c := exec.CommandContext(ctx, "kubectl", "apply", "-f", "-")
	r := bytes.NewReader([]byte(rawYAML))
	c.Stdin = r

	// TODO: handle output more cleanly
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
