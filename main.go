package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Get("/", home)
	mux.Post("/", send)
	mux.Get("/confirmation", confirmation)


	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/home.html", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
	// Step 1: Validate form
	// Step 2: Send message in an email
	// Step 3: Redirect to confirmation page
}

func confirmation(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/confirmation.html", nil)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Print(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Print(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}
