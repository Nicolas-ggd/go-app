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
		)

		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	handler := api.NewHandler(&models.UserModel{})
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
