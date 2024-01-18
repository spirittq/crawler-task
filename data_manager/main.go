package main

import (
	"datamanager/core"
	"datamanager/database"
	"datamanager/web"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	pb "shared/grpc"
	"shared/utils"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var SERVER_PORT = utils.GetEnvAsIntOrDefault("SERVER_PORT", 0)

type server struct {
	pb.CrawlerServer
}

// server method for incoming request from client. Parse, validate & save data to db.
func (s *server) CrawlerDataIncoming(stream pb.Crawler_CrawlerDataIncomingServer) error {
	var count int
	for {
		crawlerData, err := stream.Recv()
		if err == io.EOF {
			log.Info().Msgf("all data received, closing stream:: %v", count)
			success := true
			return stream.SendAndClose(&pb.CrawlerResponse{Success: &success})
		}

		log.Info().Msg("consuming crawler data")
		count++

		if err != nil {
			return err
		}
		err = core.ParseAndValidate(crawlerData)
		if err != nil {
			log.Err(err).Msg("data could not be consumed")
		}
	}
}

func main() {

	log.Info().Msg("initializing api")
	go web.App()
	log.Info().Msg("api initialized successfully")

	log.Info().Msg("initializing database")
	database.InitDB()
	log.Info().Msg("database initialized successfully")

	log.Info().Msgf("starting to listen on port %v", SERVER_PORT)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", SERVER_PORT))
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}
	log.Info().Msgf("listening on port %v", SERVER_PORT)

	log.Info().Msg("starting the server")
	s := grpc.NewServer()
	pb.RegisterCrawlerServer(s, &server{})
	log.Info().Msg("server started successfully")

	go func() {
		err = s.Serve(lis)
		if err != nil {
			log.Fatal().Msgf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info().Msg("quitting the server")
}
