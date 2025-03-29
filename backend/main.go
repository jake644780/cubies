package main

import (
	"backend/connection"
	"fmt"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	// WebSocket endpoint
	r.GET("/ws", connection.HandleWebSocket)
	
	// Start message broadcasting
	go connection.HandleMessages()

	fmt.Println("WebSocket server started on ws://localhost:8080/ws")
	r.Run(":8080")
}
