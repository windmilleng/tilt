package build

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/docker/distribution/reference"
	"github.com/opencontainers/go-digest"
	"github.com/pkg/errors"

	"github.com/windmilleng/tilt/internal/container"
	"github.com/windmilleng/tilt/internal/docker"
	"github.com/windmilleng/tilt/pkg/logger"
)

type CustomBuilder interface {
	Build(ctx context.Context, ref reference.Named, command string, expectedTag string, pushDisabled bool) (reference.NamedTagged, error)
}

type ExecCustomBuilder struct {
	dCli  docker.Client
	clock Clock
}

func NewExecCustomBuilder(dCli docker.Client, clock Clock) *ExecCustomBuilder {
	return &ExecCustomBuilder{
		dCli:  dCli,
		clock: clock,
	}
}

func (b *ExecCustomBuilder) Build(ctx context.Context, ref reference.Named, command string, expectedTag string, pushDisabled bool) (reference.NamedTagged, error) {
	if expectedTag == "" {
		expectedTag = fmt.Sprintf("tilt-build-%d", b.clock.Now().Unix())
	}

	expectedRef, err := reference.WithTag(ref, expectedTag)
	if err != nil {
		return nil, errors.Wrap(err, "CustomBuilder.Build")
	}

	cmd := exec.CommandContext(ctx, "sh", "-c", command)

	l := logger.Get(ctx)
	l.Infof("Custom Build: Injecting Environment Variables")
	l.Infof("EXPECTED_REF=%s", container.FamiliarString(expectedRef))
	env := append(os.Environ(), fmt.Sprintf("EXPECTED_REF=%s", container.FamiliarString(expectedRef)))

	for _, e := range b.dCli.Env().AsEnviron() {
		env = append(env, e)
		l.Infof("%s", e)
	}
	cmd.Env = env

	w := l.Writer(logger.InfoLvl)
	cmd.Stdout = w
	cmd.Stderr = w

	l.Infof("Running custom build cmd %q", command)
	err = cmd.Run()
	if err != nil {
		return nil, errors.Wrap(err, "Custom build command failed")
	}

	// If the command pushes the image itself, then we don't expect the image
	// to be available in the docker image registry (because the command
	// might have its own registry).
	if pushDisabled {
		return expectedRef, nil
	}

	inspect, _, err := b.dCli.ImageInspectWithRaw(ctx, expectedRef.String())
	if err != nil {
		return nil, errors.Wrap(err, "Verifying image in Docker")
	}

	dig := digest.Digest(inspect.ID)

	tag, err := digestAsTag(dig)
	if err != nil {
		return nil, errors.Wrap(err, "CustomBuilder.Build")
	}

	namedTagged, err := reference.WithTag(ref, tag)
	if err != nil {
		return nil, errors.Wrap(err, "CustomBuilder.Build")
	}

	err = b.dCli.ImageTag(ctx, dig.String(), namedTagged.String())
	if err != nil {
		return nil, errors.Wrap(err, "CustomBuilder.Build")
	}

	return namedTagged, nil
}
