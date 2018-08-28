package build

import (
	"github.com/docker/distribution/reference"

	"github.com/docker/docker/api/types/container"
	"github.com/windmilleng/tilt/internal/model"
)

type containerID string

func (cID containerID) String() string { return string(cID) }
func (cID containerID) ShortStr() string {
	if len(string(cID)) > 10 {
		return string(cID)[:10]
	}
	return string(cID)
}

// Get a container config to run a container with a given command instead of
// the existing entrypoint. If cmd is nil, we run nothing.
func containerConfigRunCmd(imgRef reference.NamedTagged, cmd model.Cmd) *container.Config {
	config := containerConfig(imgRef)

	// In Docker, both the Entrypoint and the Cmd are used to determine what
	// process the container runtime uses, where Entrypoint takes precedence over
	// command. We set both here to ensure that we don't get weird results due
	// to inheritance.
	//
	// If cmd is nil, we use a fake cmd that does nothing.
	//
	// https://github.com/opencontainers/image-spec/blob/master/config.md#properties
	if cmd.Empty() {
		config.Cmd = model.ToShellCmd("# NOTE(nick): a fake cmd").Argv
	} else {
		config.Cmd = cmd.Argv
	}
	config.Entrypoint = []string{}
	return config
}

// Get a container config to run a container as-is.
func containerConfig(imgRef reference.NamedTagged) *container.Config {
	return &container.Config{Image: imgRef.String()}
}
