package services

import (
	"fmt"
	"net/http"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")

	// Process the form data
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		// Loop through the form values and print them
		for i := 0; i <= 11; i++ {
			inputName := fmt.Sprintf("word%d", i)
			inputValue := r.FormValue(inputName)
			fmt.Printf("Input %d: %s\n", i, inputValue)
		}

		// Redirect to /home
		http.Redirect(w, r, "/home", http.StatusSeeOther)

	}
}
