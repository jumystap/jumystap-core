package handler

import "net/http"

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
    return &UserHandler{}
}

func (h *UserHandler) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
}
