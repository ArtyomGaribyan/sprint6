package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.Ldate|log.Ltime|log.Lshortfile)

	serv := server.Router(logger)

	serv.Logger.Printf("Запуск сервера")
	err := serv.Server.ListenAndServe()
	if err != nil {
		serv.Logger.Fatal(err)
	}
}
