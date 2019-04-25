package client

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/windmilleng/tilt/internal/hud/view"
	"github.com/windmilleng/tilt/internal/hud/webview"
	"github.com/windmilleng/tilt/internal/model"
	"github.com/windmilleng/tilt/internal/store"
	"github.com/windmilleng/tilt/internal/testutils/output"
)

const (
	testRoomID = model.RoomID("some-room")
	testSecret = "shh-very-secret"
)

func TestBroadcast(t *testing.T) {
	f := newFixture(t)
	defer f.TearDown()

	f.client.OnChange(f.ctx, f.store)

	f.assertNewRoomCalls(1)
	assert.Equal(t, 1, len(f.conn().json.(webview.View).Resources))
	assert.Equal(t, view.TiltfileResourceName, f.conn().json.(webview.View).Resources[0].Name.String())

	state := f.store.LockMutableStateForTesting()
	state.UpsertManifestTarget(store.NewManifestTarget(model.Manifest{Name: "fe"}))
	f.store.UnlockMutableState()

	f.client.OnChange(f.ctx, f.store)
	f.assertNewRoomCalls(1) // room already connected, shouldn't have any more NewRoom calls
	assert.Equal(t, 2, len(f.conn().json.(webview.View).Resources))
}

func TestSailRoomConnectedAction(t *testing.T) {
	f := newFixture(t)
	defer f.TearDown()
	go f.store.Loop(f.ctx)
	f.client.OnChange(f.ctx, f.store)

	time.Sleep(10 * time.Millisecond)
	for _, a := range f.getActions() {
		if roomConn, ok := a.(SailRoomConnectedAction); ok {
			assert.NoError(t, roomConn.Err)
			assert.Equal(t, "http://localhost:12345/view/some-room", roomConn.ViewURL)
			break
		}
	}
}

type fixture struct {
	t          *testing.T
	ctx        context.Context
	cancel     func()
	client     *SailClient
	store      *store.Store
	getActions func() []store.Action
}

func newFixture(t *testing.T) *fixture {
	ctx, cancel := context.WithCancel(output.CtxForTest())
	u, err := url.Parse("ws://localhost:12345")
	if err != nil {
		t.Fatal(err)
	}

	st, getActions := store.NewStoreForTesting()

	client := ProvideSailClient(model.SailURL(*u), &fakeSailRoomer{}, fakeSailDialer{})
	return &fixture{
		t:          t,
		ctx:        ctx,
		cancel:     cancel,
		client:     client,
		store:      st,
		getActions: getActions,
	}
}

func (f *fixture) conn() *fakeSailConn {
	if f.client.conn == nil {
		f.t.Fatal("client.conn is unexpectedly nil")
	}
	return f.client.conn.(*fakeSailConn)
}

func (f *fixture) assertNewRoomCalls(n int) {
	fakeRoomer, ok := f.client.roomer.(*fakeSailRoomer)
	if !ok {
		f.t.Fatal("client.roomer is not of type fakeSailRoomer??")
	}
	assert.Equal(f.t, n, fakeRoomer.newRoomCalls, "expected %d calls to NewRoom, got %d", n, fakeRoomer.newRoomCalls)
}

func (f *fixture) TearDown() {
	f.cancel()
}

type fakeSailRoomer struct {
	newRoomCalls int
}

func (r *fakeSailRoomer) NewRoom(ctx context.Context) (roomID model.RoomID, secret string, err error) {
	r.newRoomCalls += 1
	return testRoomID, testSecret, nil
}

type fakeSailDialer struct{}

func (d fakeSailDialer) DialContext(ctx context.Context, addr string, headers http.Header) (SailConn, error) {
	return &fakeSailConn{ctx: ctx}, nil
}

type fakeSailConn struct {
	ctx    context.Context
	json   interface{}
	closed bool
}

func (c *fakeSailConn) WriteJSON(v interface{}) error {
	c.json = v
	return nil
}

func (c *fakeSailConn) NextReader() (int, io.Reader, error) {
	<-c.ctx.Done()
	return 0, nil, c.ctx.Err()
}

func (c *fakeSailConn) Close() error {
	c.closed = true
	return nil
}
