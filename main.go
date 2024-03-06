package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"websocket/cmd/api"
	"websocket/cmd/api/websocket"
	"websocket/internal/db"
	"websocket/internal/models"
)

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

	database, err := db.ConnectionDB()
	if err != nil {
		panic(err)
	}
	fmt.Println(database, "<-- database")
	_ = models.NewDBWrapper(database)

	WebSocket := websocket.NewWebsocket()

	// run websocket
	go WebSocket.Run()

	srv := http.Server{
		Addr:    server.config.ListenerAddr,
		Handler: api.Routes(),
	}

	log.Printf("Server starting on %s", server.config.ListenerAddr)

	err = srv.ListenAndServe()
	log.Fatal(err)
}
