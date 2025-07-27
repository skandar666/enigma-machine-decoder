package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type MyServer struct {
	Logger *log.Logger
	Server *http.Server
}

func NewServer(logger *log.Logger) *MyServer {
	r := chi.NewRouter()

	r.Get("/", handlers.DataLoad)
	r.Post("/upload", handlers.DataReceive)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &MyServer{
		Logger: logger,
		Server: srv,
	}
}
