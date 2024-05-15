package main

import (
	v1 "ip2loc/controllers/v1"
	"ip2loc/server"
	"ip2loc/services"
	"ip2loc/startup"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	err := startup.SetupConfig()
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	server := server.NewIP2LocServer(v1.NewFindCountryController(services.NewGeoIPService()))

	err = server.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
