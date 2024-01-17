package main

import (
	"context"
	"crawler/core"
	"fmt"
	"shared/utils"

	pb "shared/grpc"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var SERVER_DOMAIN = utils.GetEnvOrDefault("SERVER_DOMAIN", "")
var SERVER_PORT = utils.GetEnvAsIntOrDefault("SERVER_PORT", 0)

func main() {

	log.Info().Msg("connecting to the server")
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v", SERVER_DOMAIN, SERVER_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Msgf("did not connect: %v", err)
	}
	log.Info().Msg("connected to the server successfully")
	defer conn.Close()

	log.Info().Msg("starting the client and establishing stream")
	c := pb.NewCrawlerClient(conn)
	stream, err := c.CrawlerDataIncoming(context.Background())
	if err != nil {
		log.Fatal().Msgf("stream failed: %v", err)
	}
	log.Info().Msg("stream established successfully")

	log.Info().Msg("starting crawler task")
	core.Crawling(stream)
	log.Info().Msg("crawler task finished, closing the stream")
	_, err = stream.CloseAndRecv()
	if err != nil {
		log.Fatal().Msgf("failed to get reply from server: %v", err)
	}
	log.Info().Msg("stream closed, exiting")
}
