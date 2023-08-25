package main

import (
	"delta_wallet/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "./static")

	router.GET("/", handlers.Index)
	router.GET("/home", handlers.Home)
	router.GET("/receive", handlers.Receive)
	router.GET("/send", handlers.Send)
	router.GET("/login", handlers.Login)
	router.GET("/create-wallet", handlers.CreateWallet)

	router.Run(":8080")
}
