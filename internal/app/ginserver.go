package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type GinServer struct{}

func NewGinServer() *GinServer {
	return &GinServer{}
}

func (g GinServer) Run(addr string) error {
	router := gin.Default()
	router.POST("/", ginShortURL)
	router.GET("/:id", ginGetURL)

	log.Printf("Starting server at %s...\r\n", addr)
	return router.Run(addr)
}
