package connect

import (
	"errors"

	"linkedin-automation-poc/internal/browser"
	"linkedin-automation-poc/internal/search"
	"linkedin-automation-poc/internal/stealth"
)

type MockConnectService struct {
	limit int
	count int
	br    *browser.Browser
	beh   stealth.Behavior
}

func NewMockConnectService(br *browser.Browser, b stealth.Behavior, limit int) *MockConnectService {
	return &MockConnectService{br: br, beh: b, limit: limit}
}

func (c *MockConnectService) Send(profile search.Profile) error {
	if c.count >= c.limit {
		return errors.New("connection limit reached")
	}
	_ = c.beh.BeforeAction()
	c.count++
	_ = c.beh.AfterAction()
	return nil
}
