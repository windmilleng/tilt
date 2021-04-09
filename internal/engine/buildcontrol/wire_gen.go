// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package buildcontrol

import (
	"context"

	"github.com/google/wire"
	"github.com/tilt-dev/wmclient/pkg/dirs"

	"github.com/tilt-dev/tilt/internal/analytics"
	"github.com/tilt-dev/tilt/internal/build"
	"github.com/tilt-dev/tilt/internal/containerupdate"
	"github.com/tilt-dev/tilt/internal/docker"
	"github.com/tilt-dev/tilt/internal/dockercompose"
	"github.com/tilt-dev/tilt/internal/dockerfile"
	"github.com/tilt-dev/tilt/internal/k8s"
	"github.com/tilt-dev/tilt/internal/tracer"
)

// Injectors from wire.go:

func ProvideImageBuildAndDeployer(ctx context.Context, docker2 docker.Client, kClient k8s.Client, env k8s.Env, dir *dirs.TiltDevDir, clock build.Clock, kp KINDLoader, analytics2 *analytics.TiltAnalytics) (*ImageBuildAndDeployer, error) {
	labels := _wireLabelsValue
	dockerImageBuilder := build.NewDockerImageBuilder(docker2, labels)
	dockerBuilder := build.DefaultDockerBuilder(dockerImageBuilder)
	execCustomBuilder := build.NewExecCustomBuilder(docker2, clock)
	updateModeFlag := _wireUpdateModeFlagValue
	runtime := k8s.ProvideContainerRuntime(ctx, kClient)
	updateMode, err := ProvideUpdateMode(updateModeFlag, env, runtime)
	if err != nil {
		return nil, err
	}
	imageBuildAndDeployer := NewImageBuildAndDeployer(dockerBuilder, execCustomBuilder, kClient, env, analytics2, updateMode, clock, runtime, kp)
	return imageBuildAndDeployer, nil
}

var (
	_wireLabelsValue         = dockerfile.Labels{}
	_wireUpdateModeFlagValue = UpdateModeFlag(UpdateModeAuto)
)

func ProvideDockerComposeBuildAndDeployer(ctx context.Context, dcCli dockercompose.DockerComposeClient, dCli docker.Client, dir *dirs.TiltDevDir) (*DockerComposeBuildAndDeployer, error) {
	labels := _wireLabelsValue
	dockerImageBuilder := build.NewDockerImageBuilder(dCli, labels)
	dockerBuilder := build.DefaultDockerBuilder(dockerImageBuilder)
	clock := build.ProvideClock()
	execCustomBuilder := build.NewExecCustomBuilder(dCli, clock)
	updateModeFlag := _wireBuildcontrolUpdateModeFlagValue
	env := _wireEnvValue
	kubeContextOverride := _wireKubeContextOverrideValue
	clientConfig := k8s.ProvideClientConfig(kubeContextOverride)
	restConfigOrError := k8s.ProvideRESTConfig(clientConfig)
	clientsetOrError := k8s.ProvideClientset(restConfigOrError)
	portForwardClient := k8s.ProvidePortForwardClient(restConfigOrError, clientsetOrError)
	namespace := k8s.ProvideConfigNamespace(clientConfig)
	config, err := k8s.ProvideKubeConfig(clientConfig, kubeContextOverride)
	if err != nil {
		return nil, err
	}
	kubeContext, err := k8s.ProvideKubeContext(config)
	if err != nil {
		return nil, err
	}
	minikubeClient := k8s.ProvideMinikubeClient(kubeContext)
	client := k8s.ProvideK8sClient(ctx, env, restConfigOrError, clientsetOrError, portForwardClient, namespace, minikubeClient, clientConfig)
	runtime := k8s.ProvideContainerRuntime(ctx, client)
	updateMode, err := ProvideUpdateMode(updateModeFlag, env, runtime)
	if err != nil {
		return nil, err
	}
	imageBuilder := NewImageBuilder(dockerBuilder, execCustomBuilder, updateMode)
	dockerComposeBuildAndDeployer := NewDockerComposeBuildAndDeployer(dcCli, dCli, imageBuilder, clock)
	return dockerComposeBuildAndDeployer, nil
}

var (
	_wireBuildcontrolUpdateModeFlagValue = UpdateModeFlag(UpdateModeAuto)
	_wireEnvValue                        = k8s.Env(k8s.EnvNone)
	_wireKubeContextOverrideValue        = k8s.KubeContextOverride("")
)

// wire.go:

var BaseWireSet = wire.NewSet(wire.Value(dockerfile.Labels{}), k8s.ProvideMinikubeClient, build.DefaultDockerBuilder, build.NewDockerImageBuilder, build.NewExecCustomBuilder, wire.Bind(new(build.CustomBuilder), new(*build.ExecCustomBuilder)), NewDockerComposeBuildAndDeployer,
	NewImageBuildAndDeployer,
	NewLocalTargetBuildAndDeployer, containerupdate.NewDockerUpdater, containerupdate.NewExecUpdater, NewImageBuilder, tracer.InitOpenTelemetry, ProvideUpdateMode,
)
