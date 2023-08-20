package main

import (
	"fmt"
	"net/http"

	"delta_wallet/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Маршруты и обработчики
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/home", handlers.HomeHandler)
	router.HandleFunc("/receive", handlers.ReceiveHandler)
	router.HandleFunc("/send", handlers.SendHandler)
	router.HandleFunc("/login", handlers.LoginHandler)

	// Добавляем маршрут для статических файлов
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", router)

	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
}
