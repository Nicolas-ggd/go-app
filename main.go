package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"websocket/cmd/api"
	"websocket/cmd/api/websocket"
	"websocket/cmd/app"
	"websocket/internal/db"
	"websocket/internal/models"
)

// Server represents a server instance with its configuration.
type Server struct {
	config *Config // Configuration for the server.
}

// NewServer creates a new Server instance with the provided configuration.
func NewServer(config *Config) (*Server, error) {
	// Validate configuration or implement error handling if necessary
	return &Server{config: config}, nil
}

// Config holds configuration options for the server.
type Config struct {
	ListenerAddr string // Address on which the server listens for connections.
}

// WithListenerAddr creates a copy of the Config with the updated ListenerAddr.
func (c *Config) WithListenerAddr(addr string) *Config {
	// Consider cloning the Config to avoid mutating the original
	return &Config{ListenerAddr: addr}
}

// NewConfig creates a new Config instance with default values.
func NewConfig() *Config {
	return &Config{
		ListenerAddr: ":7000", // Default listener address.
	}
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 167.99.246.163:7000
// @BasePath /v1

func main() {
	//addr := flag.String("addr", ":7000", "HTTP Network")

	config := NewConfig()

	server, err := NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()

	db.DBConnection()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		err := db.DB.AutoMigrate(
			&models.User{},
			&models.Token{},
		)

		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	a := app.NewApplication(websocket.NewWebsocket(), api.NewHandler())

	go a.Websocket.Run()

	srv := http.Server{
		Addr:    server.config.ListenerAddr,
		Handler: a.Routes(),
	}

	log.Printf("Server starting on %s", server.config.ListenerAddr)

	err = srv.ListenAndServe()
	log.Fatal(err)
}
