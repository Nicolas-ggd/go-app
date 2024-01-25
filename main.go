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
	addr := flag.String("addr", ":7000", "HTTP Network")

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

	handler := api.NewHandler(&models.User{}, &models.Token{})
	a := app.NewApplication(websocket.NewWebsocket(), handler)

	go a.Websocket.Run()

	srv := http.Server{
		Addr:    *addr,
		Handler: a.Routes(),
	}

	log.Printf("Server starting on %s", *addr)
	err := srv.ListenAndServe()

	log.Fatal(err)
}
