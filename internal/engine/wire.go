// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package engine

import (
	"context"

	"github.com/google/go-cloud/wire"
	"github.com/windmilleng/tilt/internal/build"
	"github.com/windmilleng/tilt/internal/k8s"
	"github.com/windmilleng/tilt/internal/synclet"
	"github.com/windmilleng/wmclient/pkg/dirs"
)

func provideBuildAndDeployer(
	ctx context.Context,
	docker build.DockerClient,
	k8s k8s.Client,
	dir *dirs.WindmillDir,
	env k8s.Env,
	sCli synclet.SyncletClient,
	shouldFallBackToImgBuild func(error) bool) (BuildAndDeployer, error) {
	wire.Build(
		// dockerImageBuilder ( = ImageBuilder)
		build.DefaultImageBuilder,
		build.NewDockerImageBuilder,
		build.DefaultConsole,
		build.DefaultOut,
		wire.Value(build.Labels{}),

		// ImageBuildAndDeployer (FallbackBuildAndDeployer)
		wire.Bind(new(FallbackBuildAndDeployer), new(ImageBuildAndDeployer)),
		NewImageBuildAndDeployer,

		// FirstLineBuildAndDeployer (LocalContainerBaD OR SyncletBaD)
		build.NewContainerUpdater, // in case it's a LocalContainerBuildAndDeployer
		NewFirstLineBuildAndDeployer,

		wire.Bind(new(BuildAndDeployer), new(CompositeBuildAndDeployer)),
		NewCompositeBuildAndDeployer,
	)

	return nil, nil
}
