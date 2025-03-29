package connection

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
)

var (
	clients   = make(map[*websocket.Conn]*models.Player) // Active clients
	usernames = make(map[string]*models.Player)
	broadcast = make(chan string)              // Channel for broadcasting messages
	upgrader  = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {return true /* Allow all origins (for testing)*/},}
	mu sync.Mutex // Mutex for safe concurrent access
)

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	// Initially, awaiting the username
	var player *models.Player

	// Read the first message to set the username
	_, msg, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("Error reading message:", err)
		return
	}

	// Unmarshal the incoming JSON message
	var setUsernameMessage models.SetUsernameMessage
	err = json.Unmarshal(msg, &setUsernameMessage)
	if err != nil {
		fmt.Println("Error unmarshalling message:", err)
		return
	}

	// Check if the username already exists
	mu.Lock()
	if existingPlayer, exists := usernames[setUsernameMessage.SetUsername]; exists {
		// Username exists, respond to the client
		conn.WriteMessage(websocket.TextMessage, []byte("Username '" + setUsernameMessage.SetUsername + "' is already taken."))
		player = existingPlayer
	} else {
		// Username doesn't exist, create a new player
		player = &models.Player{Username: setUsernameMessage.SetUsername}

		// Store player in the clients map
		clients[conn] = player
		
		// Store player's address in the usernames map
		usernames[setUsernameMessage.SetUsername] = player
		
		// Send confirmation message to the client
		conn.WriteMessage(websocket.TextMessage, []byte("Username set to '" + setUsernameMessage.SetUsername + "'"))
	}
	mu.Unlock()

	// Now, proceed with normal game logic
	fmt.Println("User", player.Username, "connected")

	// Start listening for messages after the username is set
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Client disconnected:", err)
			break
		}

		//TODO check for message type

		// Print the received message
		fmt.Printf("Received: %s\n", msg)

		// Get the player's username from the clients map
		namedMsg := fmt.Sprintf("%v said: %s", player.Username, msg)

		// Send an acknowledgment or custom response to the sender
		err = conn.WriteMessage(websocket.TextMessage, []byte("Server: You said '" + string(msg) + "'"))
		if err != nil {
			fmt.Println("Error responding to client:", err)
			break
		}

		// Broadcast the message to all clients
		broadcast <- namedMsg
	}

	// Remove disconnected client
	mu.Lock()
	delete(clients, conn)
	//TODO end game if in one and disconnected
	mu.Unlock()
}



// Broadcast messages to all connected clients
func HandleMessages() {
	for {
		msg := <-broadcast

		mu.Lock()
		for conn := range clients {
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				fmt.Println("Error sending message:", err)
				conn.Close()
				delete(clients, conn) // Remove broken connection
			}
		}
		mu.Unlock()
	}
}