package handler

import (
	"log"
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
    
    startDate := r.URL.Query().Get("start_date")
    endDate := r.URL.Query().Get("end_date")
    log.Print(endDate)
    log.Print(startDate)

    analytics, err := h.service.GetAnalytics(startDate, endDate)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
    }

    response := map[string]interface{}{
        "analytics": analytics,
    }

    utils.WriteJSON(w, http.StatusOK, response)
} 
