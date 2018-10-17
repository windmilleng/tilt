package hud

import (
	"context"

	"github.com/pkg/errors"

	"github.com/windmilleng/tilt/internal/hud/view"
	"github.com/windmilleng/tilt/internal/logger"
	"github.com/windmilleng/tilt/internal/store"
)

type HeadsUpDisplay interface {
	Run(ctx context.Context, st *store.Store) error
	Update(v view.View) error
	OnChange(ctx context.Context, st *store.Store)
	Close()
}

type Hud struct {
	r *Renderer
	a *ServerAdapter
}

var _ HeadsUpDisplay = (*Hud)(nil)

func NewDefaultHeadsUpDisplay() (HeadsUpDisplay, error) {
	return &Hud{
		r: NewRenderer(),
	}, nil
}

func (h *Hud) Run(ctx context.Context, st *store.Store) error {
	a, err := NewServer(ctx)
	if err != nil {
		return err
	}

	h.a = a

	for {
		select {
		case <-ctx.Done():
			h.Close()
			err := ctx.Err()
			if err != context.Canceled {
				return err
			} else {
				return nil
			}
		case ready := <-a.readyCh:
			err := h.r.SetUp(ready, st)
			if err != nil {
				return err
			}
		case <-a.streamClosedCh:
			h.r.Reset()
		case <-a.serverClosed:
			return nil
		}
	}
}

func (h *Hud) Close() {
	if h.a != nil {
		h.a.Close()
	}
}

func (h *Hud) OnChange(ctx context.Context, st *store.Store) {
	onChange(ctx, st, h)
}

func (h *Hud) Update(v view.View) error {
	err := h.r.Render(v)
	return errors.Wrap(err, "error rendering hud")
}

func onChange(ctx context.Context, st *store.Store, h HeadsUpDisplay) {
	state := st.RLockState()
	if len(state.ManifestStates) == 0 {
		st.RUnlockState()
		return
	}

	view := store.StateToView(state)
	st.RUnlockState()

	err := h.Update(view)
	if err != nil {
		logger.Get(ctx).Infof("Error updating HUD: %v", err)
	}
}

var _ store.Subscriber = &Hud{}
