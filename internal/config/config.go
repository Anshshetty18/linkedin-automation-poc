package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Browser BrowserConfig `yaml:"browser"`
	Limits  LimitsConfig  `yaml:"limits"`
	Timing  TimingConfig  `yaml:"timing"`
}

type BrowserConfig struct {
	Headless       bool `yaml:"headless"`
	TimeoutSeconds int  `yaml:"timeout_seconds"`
	Timeout        time.Duration
}

type LimitsConfig struct {
	DailyConnections int `yaml:"daily_connections"`
	DailyMessages    int `yaml:"daily_messages"`
}

type TimingConfig struct {
	MinDelayMs int `yaml:"min_delay_ms"`
	MaxDelayMs int `yaml:"max_delay_ms"`
}

func Load(path string) (*Config, error) {
	cfg := defaultConfig()

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	applyEnvOverrides(cfg)
	if err := validate(cfg); err != nil {
		return nil, err
	}

	cfg.Browser.Timeout = time.Duration(cfg.Browser.TimeoutSeconds) * time.Second
	return cfg, nil
}

func defaultConfig() *Config {
	return &Config{
		Browser: BrowserConfig{Headless: true, TimeoutSeconds: 30},
		Limits:  LimitsConfig{DailyConnections: 20, DailyMessages: 15},
		Timing:  TimingConfig{MinDelayMs: 300, MaxDelayMs: 1200},
	}
}

func applyEnvOverrides(cfg *Config) {
	if v := os.Getenv("BROWSER_HEADLESS"); v != "" {
		cfg.Browser.Headless = v == "true"
	}
	if v := os.Getenv("BROWSER_TIMEOUT_SECONDS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			cfg.Browser.TimeoutSeconds = n
		}
	}
}

func validate(cfg *Config) error {
	if cfg.Browser.TimeoutSeconds <= 0 {
		return errors.New("invalid browser timeout")
	}
	if cfg.Timing.MinDelayMs > cfg.Timing.MaxDelayMs {
		return errors.New("min delay > max delay")
	}
	return nil
}
