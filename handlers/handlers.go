package handlers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "layouts/base.html", gin.H{
		"Title":   "Index",
		"Content": template.HTML("<h1>Index Page</h1>"),
	})
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "layouts/base.html", gin.H{
		"Title":   "Home",
		"Content": template.HTML("<h1>Home Page</h1>"),
	})
}

func Receive(c *gin.Context) {
	c.HTML(http.StatusOK, "layouts/base.html", gin.H{
		"Title":   "Receive",
		"Content": template.HTML("<h1>Receive Page</h1>"),
	})
}

func Send(c *gin.Context) {
	c.HTML(http.StatusOK, "layouts/base.html", gin.H{
		"Title":   "Send",
		"Content": template.HTML("<h1>Send Page</h1>"),
	})
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "layouts/base.html", gin.H{
		"Title":   "Login",
		"Content": template.HTML("<h1>Login Page</h1>"),
	})
}

func CreateWallet(c *gin.Context) {
	mnemonic := "your_generated_mnemonic_here"

	c.HTML(http.StatusOK, "layouts/base.html", gin.H{
		"Title":    "Create Wallet",
		"Content":  "content/create_wallet.html",
		"Mnemonic": template.HTML(mnemonic), // Преобразуем в template.HTML
	})
}
