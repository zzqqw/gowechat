package client

import (
	"context"
	"github.com/cenkalti/backoff"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type TokenInfo struct {
	TokenStr  string        `json:"token"`
	ExpiresIn time.Duration `json:"expires_in"`
}
type Token struct {
	TokenInfo
	Mutex        *sync.RWMutex
	LastRefresh  time.Time
	GetTokenFunc func() (TokenInfo, error)
}

func (t *Token) SetGetTokenFunc(f func() (TokenInfo, error)) {
	t.GetTokenFunc = f
}

func (t *Token) GetTokenStr() string {
	// intensive mutex juggling action
	t.Mutex.RLock()
	if t.TokenStr == "" {
		// RWMutex doesn't like recursive locking
		t.Mutex.RUnlock()
		_ = t.syncToken()
		t.Mutex.RLock()
	}
	tokenToUse := t.TokenStr
	t.Mutex.RUnlock()
	return tokenToUse
}
func (t *Token) syncToken() error {
	get, err := t.GetTokenFunc()
	if err != nil {
		return err
	}
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.TokenStr = get.TokenStr
	t.ExpiresIn = get.ExpiresIn * time.Second
	t.LastRefresh = time.Now()
	return nil
}

func (t *Token) TokenRefresher(ctx context.Context) {
	// refresh per 30m
	const refreshTimeWindow = (2*60 - 30) * time.Minute
	const minRefreshDuration = 5 * time.Second
	var waitDuration time.Duration = 0
	for {
		select {
		case <-time.After(waitDuration):
			if err := backoff.Retry(t.syncToken, backoff.WithContext(backoff.NewExponentialBackOff(), ctx)); err != nil {
				logrus.Error("retry getting access toke failed err=" + err.Error())
				_ = err
			}
			waitDuration = t.LastRefresh.Add(t.ExpiresIn).Add(-refreshTimeWindow).Sub(t.LastRefresh)
			logrus.Debug("access_token", "token", t.TokenStr, "expiresIn", t.ExpiresIn, "nextRefreshTime", waitDuration)
			if waitDuration < minRefreshDuration {
				waitDuration = minRefreshDuration
			}
		case <-ctx.Done():
			return
		}
	}
}
