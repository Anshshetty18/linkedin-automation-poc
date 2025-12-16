package search

import (
	"linkedin-automation-poc/internal/browser"
	"linkedin-automation-poc/internal/stealth"
)

type MockSearchService struct {
	br       *browser.Browser
	behavior stealth.Behavior
}

func NewMockSearchService(br *browser.Browser, b stealth.Behavior) *MockSearchService {
	return &MockSearchService{br: br, behavior: b}
}

func (s *MockSearchService) FindProfiles(query string) ([]Profile, error) {
	_ = s.behavior.BeforeAction()
	_ = s.behavior.AfterAction()

	return []Profile{
		{Name: "Alex Johnson", Title: "Software Engineer", Location: "India", URL: "https://example.com/alex"},
		{Name: "Priya Sharma", Title: "Backend Developer", Location: "India", URL: "https://example.com/priya"},
	}, nil
}
