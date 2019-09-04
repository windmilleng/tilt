package cloud

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/windmilleng/tilt/internal/feature"
	"github.com/windmilleng/tilt/internal/store"
	"github.com/windmilleng/tilt/pkg/logger"
)

// to avoid infinitely resubmitting requests on error
const timeoutAfterError = 5 * time.Minute

const TiltTokenHeaderName = "X-Tilt-Token"

func NewUsernameManager(client HttpClient) *CloudUsernameManager {
	return &CloudUsernameManager{client: client}
}

type CloudUsernameManager struct {
	client HttpClient

	sleepingAfterErrorUntil time.Time
	currentlyMakingRequest  bool
	mu                      sync.Mutex
}

func ProvideHttpClient() HttpClient {
	return http.DefaultClient
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type whoAmIResponse struct {
	Found    bool
	Username string
}

func (c *CloudUsernameManager) error() {
	c.mu.Lock()
	c.sleepingAfterErrorUntil = time.Now().Add(timeoutAfterError)
	c.mu.Unlock()
}

func (c *CloudUsernameManager) CheckUsername(ctx context.Context, st store.RStore, blocking bool) {
	state := st.RLockState()
	tok := state.Token
	st.RUnlockState()

	c.mu.Lock()
	c.currentlyMakingRequest = true
	c.mu.Unlock()

	defer func() {
		c.mu.Lock()
		c.currentlyMakingRequest = false
		c.mu.Unlock()
	}()

	u := URL(state.CloudAddress)
	u.Path = "/api/whoami"

	if blocking {
		u.Query().Set("wait_for_registration", "true")
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		logger.Get(ctx).Infof("error making whoami request: %v", err)
		c.error()
		return
	}
	req.Header.Set(TiltTokenHeaderName, string(tok))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.client.Do(req)
	if err != nil {
		logger.Get(ctx).Infof("error checking tilt cloud status: %v", err)
		c.error()
		return
	}

	if resp.StatusCode != http.StatusOK {
		logger.Get(ctx).Infof("error checking tilt cloud status: %v", resp)
		c.error()
		return
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Get(ctx).Infof("error reading response body: %v", err)
		c.error()
		return
	}
	r := whoAmIResponse{}
	err = json.NewDecoder(bytes.NewReader(responseBody)).Decode(&r)
	if err != nil {
		logger.Get(ctx).Infof("error decoding tilt whoami response '%s': %v", string(responseBody), err)
		c.error()
		return
	}

	st.Dispatch(store.TiltCloudUserLookedUpAction{
		Found:    r.Found,
		Username: r.Username,
	})
}

func (c *CloudUsernameManager) OnChange(ctx context.Context, st store.RStore) {
	state := st.RLockState()
	defer st.RUnlockState()

	if !state.Features[feature.Snapshots] {
		return
	}

	c.mu.Lock()
	sleepingAfterErrorUntil := c.sleepingAfterErrorUntil
	currentlyMakingRequest := c.currentlyMakingRequest
	c.mu.Unlock()

	// if a refresh has been induced, then do a long_get lookup, so that we get the username as soon as
	// the user has finished the process
	if state.TiltCloudUsernameNeedsRefresh && !currentlyMakingRequest {
		go c.CheckUsername(ctx, st, true)
		return
	}

	// otherwise, we're not necessarily expecting a username, so just get the current state

	// c.currentlyMakingRequest is a bit of a race condition here:
	// 1. start making request that's going to return TokenKnownUnregistered = true
	// 2. before request finishes, web ui triggers refresh, setting TokenKnownUnregistered = false
	// 3. request started in (1) finishes, sets TokenKnownUnregistered = true
	// we never make a request post-(2), where the token was registered
	// This is mitigated by - a) the window between (1) and (3) is small, and b) the user can just click refresh again
	if time.Now().Before(sleepingAfterErrorUntil) ||
		state.TiltCloudUsername != "" ||
		state.TokenKnownUnregistered ||
		currentlyMakingRequest {
		return
	}

	go c.CheckUsername(ctx, st, false)
}
