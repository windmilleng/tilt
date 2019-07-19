package containerupdate

import (
	"context"
	"io"

	"github.com/windmilleng/tilt/internal/k8s"
	"github.com/windmilleng/tilt/internal/model"
	"github.com/windmilleng/tilt/internal/store"
)

type FakeContainerUpdater struct {
	ValidateErr error
	UpdateErr   error

	Calls []UpdateContainerCall
}

var _ ContainerUpdater = &FakeContainerUpdater{}

type UpdateContainerCall struct {
	DeployInfo store.DeployInfo
	Archive    io.Reader
	ToDelete   []string
	Cmds       []model.Cmd
	HotReload  bool
}

func (cu *FakeContainerUpdater) ValidateSpecs(specs []model.TargetSpec, env k8s.Env) error {
	var err error
	if cu.ValidateErr != nil {
		err = cu.ValidateErr
		cu.ValidateErr = nil
	}
	return err
}

func (cu *FakeContainerUpdater) UpdateContainer(ctx context.Context, deployInfo store.DeployInfo,
	archiveToCopy io.Reader, filesToDelete []string, cmds []model.Cmd, hotReload bool) error {
	cu.Calls = append(cu.Calls, UpdateContainerCall{
		DeployInfo: deployInfo,
		Archive:    archiveToCopy,
		ToDelete:   filesToDelete,
		Cmds:       cmds,
		HotReload:  hotReload,
	})

	var err error
	if cu.UpdateErr != nil {
		err = cu.UpdateErr
		cu.UpdateErr = nil
	}
	return err
}
