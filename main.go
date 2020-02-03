package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/CuprumBur/JourneyfromOoLanguageToGolang/api"
	"github.com/CuprumBur/JourneyfromOoLanguageToGolang/storage"
	"github.com/go-redis/redis/v7"
)

func main() {
	apiPort := flag.String("port", "8080", "service port")
	flag.Parse()

	r := redis.NewClient(&redis.Options{Addr: "localhost:" + *apiPort})
	s := storage.NewStorage(r)
	a := api.NewAPI(s)

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
		<-sigs
		log.Println("performing shutdown...")
		if err := a.Shutdown(); err != nil {
			log.Printf("failed to shutdown sever: %v", err)
		}
	}()

	log.Printf("service is ready to listen on port: %s", *apiPort)
	if err := a.Start(*apiPort); err != http.ErrServerClosed {
		log.Printf("sever failed: %v", err)
		os.Exit(1)
	}
}
