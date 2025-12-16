package message

import "linkedin-automation-poc/internal/search"

type Service interface {
	Send(profile search.Profile, message string) error
}
