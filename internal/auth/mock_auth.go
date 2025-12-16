package auth

import (
	"sync"

	"linkedin-automation-poc/internal/browser"
	"linkedin-automation-poc/internal/stealth"
)

type MockAuthenticator struct {
	once     sync.Once
	loggedIn bool
	br       *browser.Browser
	behavior stealth.Behavior
}

func NewMockAuthenticator(br *browser.Browser, b stealth.Behavior) *MockAuthenticator {
	return &MockAuthenticator{br: br, behavior: b}
}

func (a *MockAuthenticator) IsAuthenticated() (bool, error) {
	return a.loggedIn, nil
}

func (a *MockAuthenticator) Login() error {
	a.once.Do(func() {
		_ = a.behavior.BeforeAction()
		a.loggedIn = true
		_ = a.behavior.AfterAction()
	})
	return nil
}
