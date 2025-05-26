package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

var shortURLAddr string

type GinServer struct {
}

func NewGinServer(host string) *GinServer {
	shortURLAddr = host
	return &GinServer{}
}

func (g GinServer) Run(addr string) error {
	router := gin.Default()
	router.POST("/", ginShortURL)
	router.GET("/:id", ginGetURL)

	log.Printf("Starting server at %s...\r\n", addr)
	return router.Run(addr)
}
