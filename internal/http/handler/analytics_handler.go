package handler

import (
	"net/http"

	"github.com/jumystap/jumystap-core/internal/service"
	"github.com/jumystap/jumystap-core/internal/utils"
)

type AnalyticsHandler struct {
    service *service.AnalyticsService
}

func NewAnalyticsHandler(service *service.AnalyticsService) *AnalyticsHandler {
    return &AnalyticsHandler{service: service}
}

func (h *AnalyticsHandler) HandleGetAnalytics(w http.ResponseWriter, r *http.Request) {
    const op = "handler.HandleGetAnalytics"
    analytics, err := h.service.GetAnalytics()
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
    }

    response := map[string]interface{}{
        "analytics": analytics,
    }

    utils.WriteJSON(w, http.StatusOK, response)
} 
