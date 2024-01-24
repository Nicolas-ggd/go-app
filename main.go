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

	apply := &app.Application{
		Websocket: websocket.NewWebsocket(),
	}

	go apply.Websocket.Run()

	srv := http.Server{
		Addr: *addr,
	}

	log.Printf("Server starting on %s", *addr)
	err := srv.ListenAndServe()

	log.Fatal(err)
}
