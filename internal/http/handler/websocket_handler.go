package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/jumystap/jumystap-core/internal/repository"
)

type Client struct {
	conn    *websocket.Conn
	send    chan []byte
	userId  int64
}

type WebSocketHandler struct {
	upgrader    websocket.Upgrader
	clients     map[*Client]bool
	broadcast   chan []byte
	messageRepo *repository.MessageRepository
}

// NewWebSocketHandler initializes a new WebSocket handler
func NewWebSocketHandler(messageRepo *repository.MessageRepository) *WebSocketHandler {
	return &WebSocketHandler{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		clients:     make(map[*Client]bool),
		broadcast:   make(chan []byte),
		messageRepo: messageRepo,
	}
}

// HandleWebSocket upgrades the HTTP connection to a WebSocket connection
func (h *WebSocketHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    log.Printf("WebSocket connection attempt from: %s, URL: %s", r.RemoteAddr, r.URL.String())

    conn, err := h.upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Failed to upgrade WebSocket connection:", err)
        http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
        return
    }

    userIdStr := r.URL.Query().Get("user_id")
    userId, err := strconv.ParseInt(userIdStr, 10, 64)
    if err != nil || userId <= 0 {
        log.Println("Invalid user ID:", userIdStr)
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    client := &Client{conn: conn, send: make(chan []byte), userId: userId}
    h.clients[client] = true

    // Start reading and writing messages for the client
    go h.readMessages(client)
    go h.writeMessages(client)

    log.Printf("Client connected: user_id=%d", userId)
}

// readMessages continuously reads messages from the client
func (h *WebSocketHandler) readMessages(client *Client) {
	defer func() {
		client.conn.Close()
		delete(h.clients, client)
	}()

	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Unmarshal the message into a structured format
		var msgData struct {
			SenderId   int64  `json:"sender_id"`
			ReceiverId int64  `json:"receiver_id"`
			Content    string `json:"content"`
		}

		if err = json.Unmarshal(message, &msgData); err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		// Save the message to the database
		err = h.messageRepo.SaveMessage(msgData.SenderId, msgData.ReceiverId, msgData.Content, "sent", 0) // Status and resume ID can be updated as needed
		if err != nil {
			log.Println("Error saving message to database:", err)
			continue
		}

		// Broadcast the message to all clients
		h.broadcast <- message
	}
}

// writeMessages sends messages to the client
func (h *WebSocketHandler) writeMessages(client *Client) {
	for message := range client.send {
		if err := client.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Error writing message:", err)
			client.conn.Close()
			delete(h.clients, client)
			break
		}
	}
}

// StartBroadcast listens for messages to broadcast to all connected clients
func (h *WebSocketHandler) StartBroadcast() {
	for {
		message := <-h.broadcast
		for client := range h.clients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		}
	}
}

