package main

import (
	"crawler-task/core"
	"os"
	"os/signal"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/rs/zerolog/log"
)

func main() {

	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal().Msg("Could not start new scheduler")
	}

	_, err = s.NewJob(
		gocron.DurationJob(
			5*time.Hour,
		),
		gocron.NewTask(
			core.Crawling,
		),
		gocron.WithStartAt(gocron.WithStartImmediately()),
	)
	if err != nil {
		log.Fatal().Msg("Could not start new job")
	}

	s.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	err = s.Shutdown()
	if err != nil {
		log.Fatal().Msg("Could not shoutdown scheduler")
	}
}
