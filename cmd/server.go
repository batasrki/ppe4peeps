package cmd

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/batasrki/ppe4peeps/handlers"
	"github.com/batasrki/ppe4peeps/logger"
)

type Server struct {
	Port int
}

func (s *Server) ListenAndServe() error {
	router := chi.NewRouter()

	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.RequestID)
	router.Use(logger.NewStructuredLogger())
	router.Use(middleware.Recoverer)

	router.Post("/orders", handlers.NewOrder)
	router.Get("/orders", handlers.Orders)

	address := fmt.Sprintf(":%d", s.Port)
	log.WithField("address", address).Info("Server starting")

	return http.ListenAndServe(address, router)
}
