package docker

import "github.com/google/wire"

func ProvideLocalAsDefault(cli LocalClient) Client {
	return Client(cli)
}
func ProvideClusterAsDefault(cli ClusterClient) Client {
	return Client(cli)
}

// Bind a docker client that can either talk to the in-cluster
// Docker daemon or to the local Docker daemon.
var ClusterWireSet = wire.NewSet(
	ProvideClusterCli,
	ProvideLocalCli,
	ProvideLocalEnv,
	ProvideClusterEnv,
	ProvideClusterAsDefault)

// Bind a docker client that can only talk to the local Docker daemon.
var LocalWireSet = wire.NewSet(
	ProvideLocalCli,
	ProvideLocalEnv,
	ProvideLocalAsDefault)
