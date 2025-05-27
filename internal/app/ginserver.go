package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	router.POST("/", GinLogger(ginShortURL))
	router.GET("/:id", GinLogger(ginGetURL))

	log.Printf("Starting server at %s...\r\n", addr)
	return router.Run(addr)
}
