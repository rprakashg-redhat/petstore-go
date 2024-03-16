package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/rprakashg-redhat/petstore-go/internal/config"
	"github.com/rprakashg-redhat/petstore-go/pkg/http"
	"github.com/rprakashg-redhat/petstore-go/pkg/log"
)

var (
	flagConfig = flag.String("config", "./config/local.yml", "path to the config file")
	logger     log.Logger
	routes     http.Routes
)

func main() {
	flag.Parse()

	//create a root logger
	logger = log.New().With(context.TODO())

	//load application configuration
	cfg, err := config.Load(*flagConfig, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	//Set up routes for HTTP server
	routes = http.Routes{}

	// build HTTP server
	address := fmt.Sprintf(":%v", cfg.ServerPort)

	//Initialize the server
	s := http.NewServer(address, cfg.TLS, cfg.CertFile, cfg.KeyFile, routes)
	//Start the server
	s.Start(logger)
}
