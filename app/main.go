package main

import (
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/thundercats-after-dark/introspection-api/app/config"
)

const defaultPort = "8080"

func main() {
	var cfg config.Config
	envconfig.MustProcess("", &cfg)

	router, err := Router(&cfg)
	if err != nil {
		log.Fatalf("Failed to build router: %v", err)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, router))
}
