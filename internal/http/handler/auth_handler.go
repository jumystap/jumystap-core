package handler

import (
	"net/http"

	"github.com/jumystap/jumystap-core/internal/model"
	"github.com/jumystap/jumystap-core/internal/service"
	"github.com/jumystap/jumystap-core/internal/utils"
)

type AuthHandler struct {
    service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
    return &AuthHandler{service: service}
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
    var loginRequest *model.LoginRequest  
    if err := utils.ParseJSON(r, &loginRequest); err != nil {
        utils.WriteError(w, http.StatusBadGateway, err)
        return
    }
    
    user, err := h.service.Login(loginRequest.Email, loginRequest.Password)
    if err != nil {
        utils.WriteError(w, http.StatusBadGateway, err)
        return
    }

    response := map[string]interface{}{
        "user": user,
    }
    
    utils.WriteJSON(w, http.StatusOK, response)
    return
}

func (h *AuthHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
    return
}
