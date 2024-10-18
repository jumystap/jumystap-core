package model

import "time"

type Chat struct {
    PartnerId   int64     `json:"partner_id"`
    PartnerName string    `json:"partner_name"` 
    LastMessage string    `json:"last_message"`  
    CreatedAt   time.Time `json:"created_at"`    
    Status      string    `json:"status"` 
}

type Message struct {
    Id          int64       `json:"id"`
    SenderId    int64       `json:"sender_id"`
    ReceiverId  int64       `json:"receiver_id"`
    Content     string      `json:"content"`
    ResumeId    int64       `json:"user_resume_id"`
    Status      string      `json:"status"`
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
}
