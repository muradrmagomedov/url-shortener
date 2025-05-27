package main

import (
	"fmt"
	"log"

	server "github.com/muradrmagomedov/url-shortener/internal/app"
)

const (
	host            = "localhost"
	port            = "8080"
	letters         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	shortenerLength = 8
)

func main() {
	addr := createAddr(host, port)
	// server := server.NewServer()
	server := server.NewGinServer()
	err := server.Run(addr)
	if err != nil {
		log.Fatal(err)
	}

}

func createAddr(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}
