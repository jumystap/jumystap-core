package handler

import (
	"net/http"

	"github.com/jumystap/jumystap-core/internal/service"
	"github.com/jumystap/jumystap-core/internal/utils"
)

type AnnouncementHandler struct {
    service *service.AnnouncementService
}

func NewAnnouncementHandler (service *service.AnnouncementService) *AnnouncementHandler {
    return &AnnouncementHandler{service: service}
}

func (h *AnnouncementHandler) HandleGetAllAnnouncements (w http.ResponseWriter, r *http.Request) {
    announcements, err := h.service.GetAllAnnouncements()

    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    result := map[string]interface{}{
        "count":        len(announcements),  
		"announcements": announcements,       
    }

    utils.WriteJSON(w, http.StatusOK, result) 

    return
}
