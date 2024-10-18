package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jumystap/jumystap-core/internal/repository"
)

type ChatHandler struct {
	repository *repository.MessageRepository
}

func NewChatHandler(repo *repository.MessageRepository) *ChatHandler {
	return &ChatHandler{repository: repo}
}

func (h *ChatHandler) HandleGetChats(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid sender ID", http.StatusBadRequest)
		return
	}

	chats, err := h.repository.GetChats(userId)
	if err != nil {
		http.Error(w, "Failed to retrieve chats", http.StatusInternalServerError)
		return
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK) 
    json.NewEncoder(w).Encode(chats)
}

func (h *ChatHandler) HandleGetMessages(w http.ResponseWriter, r *http.Request) {
	senderIdStr := r.URL.Query().Get("sender_id")
	receiverIdStr := r.URL.Query().Get("receiver_id")

	senderId, err := strconv.ParseInt(senderIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid sender ID", http.StatusBadRequest)
		return
	}

	receiverId, err := strconv.ParseInt(receiverIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid receiver ID", http.StatusBadRequest)
		return
	}

	messages, err := h.repository.GetMessages(senderId, receiverId)
	if err != nil {
        log.Println("Failed to retrieve messages:", err)
		http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

