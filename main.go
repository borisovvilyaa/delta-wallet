package main

import (
	"delta_wallet/handlers"
	"delta_wallet/services"
	"fmt"
	"net/http"

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
	handlers.SetupRoutes()                                // Add this line
	router.HandleFunc("/sign-in", services.SignInHandler) // Use "/sign-in" as the route

	// Добавляем маршрут для статических файлов
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", router)

	fmt.Println("http://localhost:8080/login")
	http.ListenAndServe(":8080", nil)
}
