package stealth

import (
	"math/rand"
	"time"

	"linkedin-automation-poc/internal/config"
)

type SimulatedBehavior struct {
	min time.Duration
	max time.Duration
}

func NewSimulatedBehavior(cfg config.TimingConfig) *SimulatedBehavior {
	return &SimulatedBehavior{
		min: time.Duration(cfg.MinDelayMs) * time.Millisecond,
		max: time.Duration(cfg.MaxDelayMs) * time.Millisecond,
	}
}

func (s *SimulatedBehavior) delay() {
	time.Sleep(s.min + time.Duration(rand.Int63n(int64(s.max-s.min))))
}

func (s *SimulatedBehavior) BeforeAction() error {
	s.delay()
	return nil
}

func (s *SimulatedBehavior) AfterAction() error {
	s.delay()
	return nil
}
