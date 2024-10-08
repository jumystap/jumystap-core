package jumystap

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jumystap/jumystap-core/internal/http/handler"
	"github.com/jumystap/jumystap-core/internal/repository"
	"github.com/jumystap/jumystap-core/internal/service"
)

// APIServer represents the API server
type APIServer struct {
	address string
	db      *sql.DB
}

// NewAPIServer creates a new APIServer instance
func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

// CORS middleware function to handle CORS requests
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Adjust as needed for security
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Run starts the API server
func (s *APIServer) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	router.Use(corsMiddleware) // Add CORS middleware here

	analyticsRepository := repository.NewAnalyticsRepository(s.db)
	analyticsService := service.NewAnalyticsService(analyticsRepository)
	analyticsHandler := handler.NewAnalyticsHandler(analyticsService)

	router.Route("/api/v1", func(router chi.Router) {
		router.Route("/analytics", func(router chi.Router) {
			router.Get("/", analyticsHandler.HandleGetAnalytics)
		})
	})

	log.Print("Listening on", s.address)
	return http.ListenAndServe(s.address, router)
}

