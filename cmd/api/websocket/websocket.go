package websocket

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Event string `json:"event"`
	Data  any
}

type Websocket struct {
	Clients map[*Client]bool

	Broadcast chan []byte

	Register chan *Client

	Unregister chan *Client
}

func NewWebsocket() *Websocket {
	return &Websocket{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (ws *Websocket) Run() {

	for {
		select {
		case client := <-ws.Register:
			ws.Clients[client] = true

			packet := Message{
				Event: "symbols",
				Data:  "Hello from socket",
			}

			symbolByte, _ := json.Marshal(packet)
			client.send <- symbolByte

		case client := <-ws.Unregister:
			fmt.Println("unregistering client")
			if _, ok := ws.Clients[client]; ok {
				delete(ws.Clients, client)
				close(client.send)
			}
		case message := <-ws.Broadcast:
			for client := range ws.Clients {
				select {
				case client.send <- message:
					fmt.Println("broadcasting client.send ")
				default:
					close(client.send)
					delete(ws.Clients, client)
				}
			}
		}

	}
}
