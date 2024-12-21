package main

import (
	"log"
	"net/http"

	"github.com/TRXLHEAD/rpn/internal/application"
)

func main() {
	http.HandleFunc("/api/v1/calculate", application.CalcHandler)
	log.Printf("Сервер запущен")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ошибка при запуске сервера", err)
	}
}
