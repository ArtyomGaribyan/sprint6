package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Serv struct {
	Logger *log.Logger
	Server *http.Server
}

func Router(log *log.Logger) *Serv {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleMain)
	mux.HandleFunc("/upload", handlers.HandleUpload)

	serv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     log,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 15,
	}

	return &Serv{log, serv}
}
