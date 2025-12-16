package message

import (
	"errors"

	"linkedin-automation-poc/internal/browser"
	"linkedin-automation-poc/internal/search"
	"linkedin-automation-poc/internal/stealth"
)

type MockMessageService struct {
	limit int
	count int
	br    *browser.Browser
	beh   stealth.Behavior
}

func NewMockMessageService(br *browser.Browser, b stealth.Behavior, limit int) *MockMessageService {
	return &MockMessageService{br: br, beh: b, limit: limit}
}

func (m *MockMessageService) Send(profile search.Profile, msg string) error {
	if m.count >= m.limit {
		return errors.New("message limit reached")
	}
	_ = m.beh.BeforeAction()
	m.count++
	_ = m.beh.AfterAction()
	return nil
}
