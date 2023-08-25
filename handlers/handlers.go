package handlers

import (
	"delta_wallet/services"
	"html/template"
	"net/http"
)

var templateFuncs = template.FuncMap{
	"until": func(count int) []struct{} {
		return make([]struct{}, count)
	},
}

func renderTemplate(w http.ResponseWriter, tmplName string, title string, r *http.Request) {
	tmpl := template.Must(template.New("").Funcs(templateFuncs).ParseFiles(
		"templates/layouts/base.html",
		"templates/layouts/header.html",
		"templates/pages/"+tmplName,
	))
	// sessionValue, err := services.GetCryptoSessionValue(r, "address")
	// if err != nil {
	// 	fmt.Println("Error getting session value:", err)
	// }

	// // Perform a type assertion to convert sessionValue to string
	// address_session, ok := sessionValue.(string)
	// if !ok {
	// 	fmt.Println("Error: sessionValue is not a string")
	// 	return
	// }
	data := struct {
		Title string
	}{
		Title: title,
	}
	tmpl.ExecuteTemplate(w, "base.html", data)

}

// tb1qde7kxlnvau5824tq9jt5gy3xfkv3563llh70rt
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", "Home Page", r)
}

func ReceiveHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "receive.html", "Receive Page", r)
}

func SendHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "send.html", "Send Page", r)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html", "Login Page", r)
}

func SetupRoutes() {
	// Your existing route handlers

	http.HandleFunc("/sign-in", services.SignInHandler)
}
