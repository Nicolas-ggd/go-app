package main

import (
	"websocket/internal/db"
	"websocket/pkg/http/routes"
)

func main() {
	database := db.ConnectionDB()
	defer database.Close()

	routes.ServeApp()
}
