package jumystap

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jumystap/jumystap-core/internal/http/handler"
)

type APIServer struct {
    address string
    db *sql.DB
}

func NewAPIServer (address string, db *sql.DB) *APIServer {
    return &APIServer{
        address: address,
        db: db,
    }
}

func (s *APIServer) Run() error {
    router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)     

    userHandler := handler.NewUserHandler()

    router.Route("/api/v1", func(router chi.Router) { 
        router.Route("/users", func(router chi.Router) {
            router.Get("/", userHandler.HandleGetAllUsers)
        })
    })

    log.Print("Listening on", s.address)
    return http.ListenAndServe(s.address, router)
}
