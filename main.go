package main

import (
	"flag"
	"log"
	"net/http"
	"websocket/cmd/api/websocket"
	"websocket/cmd/app"
)

func main() {
	addr := flag.String("addr", ":7000", "HTTP Network")

	flag.Parse()

	a := &app.Application{
		Websocket: websocket.NewWebsocket(),
	}
	go a.Websocket.Run()

	srv := http.Server{
		Addr:    *addr,
		Handler: a.Routes(),
	}

	log.Printf("Server starting on %s", *addr)
	err := srv.ListenAndServe()

	log.Fatal(err)
}
