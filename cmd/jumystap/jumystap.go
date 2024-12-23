package jumystap

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jumystap/jumystap-core/internal/http/handler"
	local_middleware  "github.com/jumystap/jumystap-core/internal/http/middleware"
	"github.com/jumystap/jumystap-core/internal/repository"
	"github.com/jumystap/jumystap-core/internal/service"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (s *APIServer) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	router.Use(local_middleware.CorsMiddleware)

	analyticsRepository := repository.NewAnalyticsRepository(s.db)
	analyticsService := service.NewAnalyticsService(analyticsRepository)
	analyticsHandler := handler.NewAnalyticsHandler(analyticsService)

    authRepository := repository.NewAuthRepository(s.db)
    authService := service.NewAuthService(authRepository)
    authHandler := handler.NewAuthHandler(authService)

    announcementRepository := repository.NewAnnouncementRepository(s.db)
    announcementService := service.NewAnnouncementService(announcementRepository)
    announcementHandler := handler.NewAnnouncementHandler(announcementService)

    messageRepository := repository.NewMessageRepository(s.db)
    webSocketHandler := handler.NewWebSocketHandler(messageRepository)
    chatHandler := handler.NewChatHandler(messageRepository)

    go webSocketHandler.StartBroadcast()

	router.Route("/api/v1", func(router chi.Router) {
		router.Route("/analytics", func(router chi.Router) {
			router.Get("/", analyticsHandler.HandleGetAnalytics)
		})
        router.Route("/announcement", func(router chi.Router) {
			router.Get("/", announcementHandler.HandleGetAllAnnouncements)
		})
        router.Post("/login", authHandler.HandleLogin)
        router.Post("/register", authHandler.HandleRegister)

        router.Get("/ws", webSocketHandler.HandleWebSocket)
        router.Get("/chats", chatHandler.HandleGetChats) // Get chats
	    router.Get("/messages", chatHandler.HandleGetMessages)
	})

	log.Print("Listening on", s.address)
	return http.ListenAndServe(s.address, router)
}

