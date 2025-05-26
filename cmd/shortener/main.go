package main

import (
	"log"

	server "github.com/muradrmagomedov/url-shortener/internal/app"
	"github.com/muradrmagomedov/url-shortener/internal/config"
)

func main() {
	ctf := config.NewConfig()
	// server := server.NewServer()
	server := server.NewGinServer()
	err := server.Run(ctf.Addr)
	if err != nil {
		log.Fatal(err)
	}

}

// func createAddr(host, port string) string {
// 	return fmt.Sprintf("%s:%s", host, port)
// }
