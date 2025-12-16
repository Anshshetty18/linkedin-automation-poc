package main

import (
	"linkedin-automation-poc/internal/auth"
	"linkedin-automation-poc/internal/browser"
	"linkedin-automation-poc/internal/config"
	"linkedin-automation-poc/internal/connect"
	"linkedin-automation-poc/internal/message"
	"linkedin-automation-poc/internal/search"
	"linkedin-automation-poc/internal/stealth"
	"linkedin-automation-poc/internal/store"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	log.Info("starting automation proof-of-concept")

	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.WithError(err).Fatal("config load failed")
	}

	br, err := browser.New(cfg.Browser)
	if err != nil {
		log.WithError(err).Fatal("browser init failed")
	}
	defer br.Close()

	behavior := stealth.NewSimulatedBehavior(cfg.Timing)

	db, err := store.New("automation.db")
	if err != nil {
		log.WithError(err).Fatal("store init failed")
	}

	authenticator := auth.NewMockAuthenticator(br, behavior)
	if ok, _ := authenticator.IsAuthenticated(); !ok {
		if err := authenticator.Login(); err != nil {
			log.WithError(err).Fatal("login failed")
		}
	}

	searchSvc := search.NewMockSearchService(br, behavior)
	connectSvc := connect.NewMockConnectService(br, behavior, cfg.Limits.DailyConnections)
	messageSvc := message.NewMockMessageService(br, behavior, cfg.Limits.DailyMessages)

	profiles, err := searchSvc.FindProfiles("software engineer")
	if err != nil {
		log.WithError(err).Error("search failed")
		return
	}

	for _, profile := range profiles {
		log.WithField("profile", profile.URL).Info("processing profile")

		if exists, _ := db.HasConnection(profile.URL); !exists {
			if err := connectSvc.Send(profile); err == nil {
				_ = db.SaveConnection(profile.URL, "sent")
			}
		}

		if exists, _ := db.HasMessage(profile.URL); !exists {
			if err := messageSvc.Send(profile, "Hi, great connecting with you!"); err == nil {
				_ = db.SaveMessage(profile.URL, "intro message")
			}
		}
	}

	log.Info("automation run completed")
}
