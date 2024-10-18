package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jumystap/jumystap-core/internal/model"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
    return &MessageRepository{db: db}
}

func (r *MessageRepository) SaveMessage(senderId, receiverId int64, content, status string, resumeId int64) error {
	query := `
		INSERT INTO messages (sender_id, receiver_id, content, user_resume_id, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	createdAt := time.Now()
	_, err := r.db.Exec(query, senderId, receiverId, content, resumeId, status, createdAt, createdAt)
	if err != nil {
		return fmt.Errorf("could not save message: %w", err)
	}

	return nil
}

func (r *MessageRepository) GetMessages(senderId, receiverId int64) ([]model.Message, error) {
	query := `
		SELECT id, sender_id, receiver_id, content, user_resume_id, status, created_at, updated_at 
		FROM messages 
		WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)
		ORDER BY created_at
	`

	rows, err := r.db.Query(query, senderId, receiverId, receiverId, senderId)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve messages: %w", err)
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		var msg model.Message
		if err := rows.Scan(&msg.Id, &msg.SenderId, &msg.ReceiverId, &msg.Content, &msg.ResumeId, &msg.Status, &msg.CreatedAt, &msg.UpdatedAt); err != nil {
			return nil, fmt.Errorf("could not scan message: %w", err)
		}
		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while reading rows: %w", err)
	}

	return messages, nil
}

func (r *MessageRepository) GetChats(userId int64) ([]model.Chat, error) {
    query := `
        SELECT 
            chat_partner_id,
            u_partner.name AS chat_partner_name,
            m.content AS last_message,
            m.created_at,
            m.status
        FROM messages m
        JOIN users u_partner ON 
            (CASE 
                WHEN m.sender_id = ? THEN m.receiver_id
                ELSE m.sender_id
            END) = u_partner.id
        JOIN (
            SELECT 
                CASE 
                    WHEN sender_id = ? THEN receiver_id
                    ELSE sender_id
                END AS chat_partner_id,
                MAX(id) AS last_message_id
            FROM messages
            WHERE sender_id = ? OR receiver_id = ?
            GROUP BY chat_partner_id
        ) latest ON m.id = latest.last_message_id
        ORDER BY m.created_at DESC
    `

    rows, err := r.db.Query(query, userId, userId, userId, userId)
    if err != nil {
        log.Println("Error retrieving chats:", err)
        return nil, err
    }
    defer rows.Close()

    var chats []model.Chat
    for rows.Next() {
        var chat model.Chat
        if err := rows.Scan(&chat.PartnerId, &chat.PartnerName, &chat.LastMessage, &chat.CreatedAt, &chat.Status); err != nil {
            log.Println("Error scanning chat:", err)
            return nil, err
        }
        chats = append(chats, chat)
    }

    return chats, nil
}

