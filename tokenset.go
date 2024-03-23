package apitoken

import (
	"sync"
	"time"
)

type TokenSet interface {
	Add(string, time.Time)
	Validate(string) bool
	Remove(string)
	New() string
}

type NeverExpireTokens struct {
	tokens    map[string]bool
	mu        sync.RWMutex
	tokenSize int
}

func (ts *NeverExpireTokens) Add(token string, _ time.Time) {
	ts.add(token)
}

func (ts *NeverExpireTokens) add(token string) {
	ts.mu.Lock()
	ts.tokens[token] = true
	ts.mu.Unlock()
}

func (ts *NeverExpireTokens) Validate(token string) bool {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	return ts.tokens[token]
}

func (ts *NeverExpireTokens) Remove(token string) {
	ts.mu.Lock()
	delete(ts.tokens, token)
	ts.mu.Unlock()
}

func (ts *NeverExpireTokens) New() string {
	token := randomString(ts.tokenSize)
	ts.add(token)
	return token
}

func (ts *NeverExpireTokens) Set(tokens []string) {
	tokensMap := make(map[string]bool, len(tokens))

	for _, token := range tokens {
		tokensMap[token] = true
	}

	ts.mu.Lock()
	ts.tokens = tokensMap
	ts.mu.Unlock()
}

type Tokens struct {
	tokens map[string]time.Time
	mu     sync.RWMutex

	tokenSize  int
	expiration time.Duration
}

func (t *Tokens) Add(token string, expire time.Time) {
	t.add(token, expire)
}

func (t *Tokens) add(token string, expire time.Time) {
	t.mu.Lock()
	t.tokens[token] = expire
	t.mu.Unlock()
}

func (t *Tokens) Validate(token string) bool {
	t.mu.RLock()
	expire, ok := t.tokens[token]
	t.mu.RUnlock()
	if !ok {
		return false
	}
	if expire.Before(time.Now()) {
		return false
	}
	return true
}

func (t *Tokens) Remove(token string) {
	t.mu.Lock()
	delete(t.tokens, token)
	t.mu.Unlock()
}

func (t *Tokens) New() string {
	token := randomString(t.tokenSize)
	expire := time.Now().Add(t.expiration)

	t.add(token, expire)

	return token
}
