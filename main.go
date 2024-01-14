package main

import (
	"crawler-task/core"
	"crawler-task/database"
	"crawler-task/utils"
	"os"
	"os/signal"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/rs/zerolog/log"
)

func main() {

	database.InitDB()
	defer database.Database.Close()

	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal().Msg("Could not start new scheduler")
	}

	_, err = s.NewJob(
		gocron.DurationJob(
			time.Duration(utils.GetEnvAsIntOrDefault("INTERVAL_SECONDS", 120))*time.Second,
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
