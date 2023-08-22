package handlers

import (
	"delta_wallet/services"
	"html/template"
	"net/http"
)

type PageData struct {
	Title string
}

var templateFuncs = template.FuncMap{
	"until": func(count int) []struct{} {
		return make([]struct{}, count)
	},
}

func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	tmpl := template.Must(template.New("").Funcs(templateFuncs).ParseFiles(
		"templates/layouts/base.html",
		"templates/layouts/header.html",
		"templates/pages/"+tmplName,
	))

	tmpl.ExecuteTemplate(w, "base.html", data)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Home Page",
	}
	renderTemplate(w, "index.html", data)
}

func ReceiveHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Receive Page",
	}
	renderTemplate(w, "receive.html", data)
}

func SendHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Send Page",
	}
	renderTemplate(w, "send.html", data)
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Login Page",
	}
	renderTemplate(w, "login.html", data)
}

func SetupRoutes() {
	// Your existing route handlers

	http.HandleFunc("/sign-in", services.SignInHandler)
}
