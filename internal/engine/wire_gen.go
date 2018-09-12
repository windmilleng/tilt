// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package engine

import (
	context "context"
	wire "github.com/google/go-cloud/wire"
	build "github.com/windmilleng/tilt/internal/build"
	k8s "github.com/windmilleng/tilt/internal/k8s"
	synclet "github.com/windmilleng/tilt/internal/synclet"
	analytics "github.com/windmilleng/wmclient/pkg/analytics"
	dirs "github.com/windmilleng/wmclient/pkg/dirs"
)

// Injectors from wire.go:

func provideBuildAndDeployer(ctx context.Context, docker build.DockerClient, k8s2 k8s.Client, dir *dirs.WindmillDir, env k8s.Env, sCli synclet.SyncletClient, shouldFallBackToImgBuild FallbackTester) (BuildAndDeployer, error) {
	syncletBuildAndDeployer := NewSyncletBuildAndDeployer(sCli, k8s2)
	containerUpdater := build.NewContainerUpdater(docker)
	memoryAnalytics := analytics.NewMemoryAnalytics()
	localContainerBuildAndDeployer := NewLocalContainerBuildAndDeployer(containerUpdater, env, k8s2, memoryAnalytics)
	console := build.DefaultConsole()
	writer := build.DefaultOut()
	labels := _wireLabelsValue
	dockerImageBuilder := build.NewDockerImageBuilder(docker, console, writer, labels)
	imageBuilder := build.DefaultImageBuilder(dockerImageBuilder)
	imageBuildAndDeployer := NewImageBuildAndDeployer(imageBuilder, k8s2, env, memoryAnalytics)
	buildOrder := DefaultBuildOrder(syncletBuildAndDeployer, localContainerBuildAndDeployer, imageBuildAndDeployer, env)
	compositeBuildAndDeployer := NewCompositeBuildAndDeployer(buildOrder, shouldFallBackToImgBuild)
	return compositeBuildAndDeployer, nil
}

var (
	_wireLabelsValue = build.Labels{}
)

// wire.go:

var DeployerWireSet = wire.NewSet(build.DefaultConsole, build.DefaultOut, wire.Value(build.Labels{}), build.DefaultImageBuilder, build.NewDockerImageBuilder, NewImageBuildAndDeployer, build.NewContainerUpdater, NewSyncletBuildAndDeployer,
	NewLocalContainerBuildAndDeployer,
	DefaultBuildOrder, wire.Bind(new(BuildAndDeployer), new(CompositeBuildAndDeployer)), NewCompositeBuildAndDeployer)
