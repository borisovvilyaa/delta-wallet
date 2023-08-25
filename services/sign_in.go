package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignInHandler(c *gin.Context) {
	fmt.Println("hello")

	// Process the form data
	if c.Request.Method == http.MethodPost {
		err := c.Request.ParseForm()
		if err != nil {
			c.String(http.StatusBadRequest, "Error parsing form data")
			return
		}

		// Loop through the form values and print them
		for i := 0; i <= 11; i++ {
			inputName := fmt.Sprintf("word%d", i)
			inputValue := c.PostForm(inputName)
			fmt.Printf("Input %d: %s\n", i, inputValue)
		}

		// Redirect to /home
		c.Redirect(http.StatusSeeOther, "/home")
	}
}
