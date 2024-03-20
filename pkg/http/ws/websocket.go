package ws

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Event string `json:"event"`
	Data  any
}

type Websocket struct {
	Clients map[string]*Client

	Broadcast chan []byte

	Register chan *Client

	Unregister chan *Client
}

func NewWebsocket() *Websocket {
	return &Websocket{
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (ws *Websocket) Run() {
	for {
		select {
		// handle register client case
		case client := <-ws.Register:
			ws.Clients[client.ClientId] = client
			fmt.Println(client.ClientId, "<--- Register client id")
			// send default message from ws
			packet := Message{
				Event: "symbols",
				Data:  "Hello from socket",
			}

			// marshal packet and send in to the channel
			symbolByte, _ := json.Marshal(packet)
			client.Send <- symbolByte

		// unregister client case
		case client := <-ws.Unregister:
			if _, ok := ws.Clients[client.ClientId]; ok {
				// delete client
				delete(ws.Clients, client.ClientId)
				close(client.Send)
			}

		// handle case to receiving broadcast
		case message := <-ws.Broadcast:
			for _, client := range ws.Clients {
				select {
				case client.Send <- message:
					fmt.Println("broadcasting client.send ")
				default:
					close(client.Send)
					delete(ws.Clients, client.ClientId)
				}
			}
		}

	}
}
