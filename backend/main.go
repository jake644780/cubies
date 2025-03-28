package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool) // Active clients
	broadcast = make(chan string)              // Channel for broadcasting messages
	upgrader  = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {return true /* Allow all origins (for testing)*/},}
	mu sync.Mutex // Mutex for safe concurrent access
)

// WebSocket handler
func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	// Add new client
	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	fmt.Println("New client connected")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Client disconnected:", err)
			break
		}

		fmt.Printf("Received: %s\n", msg)

		// Send an acknowledgment or custom response to the sender
		response := fmt.Sprintf("Server: You said '%s'", msg)
		err = conn.WriteMessage(websocket.TextMessage, []byte(response))
		if err != nil {
			fmt.Println("Error responding to client:", err)
			break
		}

		// Broadcast the message to all clients
		broadcast <- string(msg)
	}

	// Remove disconnected client
	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
}

// Broadcast messages to all connected clients
func handleMessages() {
	for {
		msg := <-broadcast

		mu.Lock()
		for conn := range clients {
			err := conn.WriteMessage(websocket.TextMessage, []byte("Broadcast: "+msg))
			if err != nil {
				fmt.Println("Error sending message:", err)
				conn.Close()
				delete(clients, conn) // Remove broken connection
			}
		}
		mu.Unlock()
	}
}

func main() {
	r := gin.Default()

	// WebSocket endpoint
	r.GET("/ws", func(c *gin.Context) {
		handleWebSocket(c)
	})

	// Start message broadcasting
	go handleMessages()

	fmt.Println("WebSocket server started on ws://localhost:8080/ws")
	r.Run(":8080")
}
