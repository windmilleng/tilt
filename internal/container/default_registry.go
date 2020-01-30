package container

import (
	"fmt"
	"regexp"

	"github.com/docker/distribution/reference"
	"github.com/pkg/errors"
)

var escapeRegex = regexp.MustCompile(`[/:@]`)

func escapeName(s string) string {
	return string(escapeRegex.ReplaceAll([]byte(s), []byte("_")))
}

type Registry struct {
	// The Host of a container registry where we can push images. e.g.:
	//   - localhost:32000
	//   - gcr.io/windmill-public-containers
	Host string

	// The prefix we use with image names when referring to them from inside the cluster.
	// In most cases, this is equivalent to Host (the host of the container registry that we push to),
	// but sometimes users will specify a hostFromCluster separately (e.g. using a local registry with KIND:
	// YAMLs will specify the image as `registry:5000/my-img`, so the hostFromCluster will be `registry:5000`).
	hostFromCluster string
}

func (r Registry) Empty() bool { return r.Host == "" }

func NewRegistry(host string) Registry {
	// TODO(maia): validate
	return Registry{Host: host}
}

func NewRegistryWithHostFromCluster(host, fromCluster string) Registry {
	// TODO(maia): validate
	return Registry{Host: host, hostFromCluster: fromCluster}
}

// HostFromCluster returns the registry to be used from within the k8s cluster
// (e.g. in k8s YAML). Returns hostFromCluster, if specified; otherwise the Host.
func (r Registry) HostFromCluster() string {
	if r.hostFromCluster != "" {
		return r.hostFromCluster
	}
	return r.Host
}

// replaceRegistry produces a new image name that is in the specified registry.
// The name might be ugly, favoring uniqueness and simplicity and assuming that the prettiness of ephemeral dev image
// names is not that important.
func replaceRegistry(defaultReg string, rs RefSelector) (reference.Named, error) {
	if defaultReg == "" {
		return rs.AsNamedOnly(), nil
	}

	// validate the ref produced
	newNs := fmt.Sprintf("%s/%s", defaultReg, escapeName(rs.RefFamiliarName()))
	newN, err := reference.ParseNamed(newNs)
	if err != nil {
		return nil, errors.Wrapf(err, "Error parsing %s after applying default registry %s", newNs, defaultReg)
	}

	return newN, nil
}

func (r Registry) ReplaceRegistryForLocalRef(rs RefSelector) (reference.Named, error) {
	return replaceRegistry(r.Host, rs)
}

func (r Registry) ReplaceRegistryForClusterRef(rs RefSelector) (reference.Named, error) {
	return replaceRegistry(r.HostFromCluster(), rs)
}
