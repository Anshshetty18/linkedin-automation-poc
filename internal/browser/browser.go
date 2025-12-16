package browser

import (
	"context"

	"linkedin-automation-poc/internal/config"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Browser struct {
	rod    *rod.Browser
	ctx    context.Context
	cancel context.CancelFunc
}

func New(cfg config.BrowserConfig) (*Browser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)

	u := launcher.New().Headless(cfg.Headless).MustLaunch()
	b := rod.New().ControlURL(u).Context(ctx)

	if err := b.Connect(); err != nil {
		cancel()
		return nil, err
	}

	return &Browser{rod: b, ctx: ctx, cancel: cancel}, nil
}

func (b *Browser) Close() {
	b.cancel()
	_ = b.rod.Close()
}
