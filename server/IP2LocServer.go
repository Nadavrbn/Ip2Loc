package server

import (
	v1 "ip2loc/controllers/v1"

	"github.com/gin-gonic/gin"
)

type IP2LocServer struct {
	FindCountryController v1.IFindCountryController
	engine                *gin.Engine
}

func NewIP2LocServer(findCountryController v1.IFindCountryController) *IP2LocServer {
	engine := gin.Default()
	engine.Use(RateLimitingMiddleware())

	server := &IP2LocServer{
		FindCountryController: findCountryController,
		engine:                engine,
	}

	engine.GET("/v1/find-country", server.FindCountryController.FindCountry)

	return server
}

func (s *IP2LocServer) Run() error {
	return s.engine.Run()
}
