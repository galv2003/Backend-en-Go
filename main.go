package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type PageVariables struct {
	Title   string
	Message string
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/submit", SubmitForm)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Title:   "Checklist por Gabriel Alvarez",
		Message: "",
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SubmitForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		message := r.FormValue("message")
		fmt.Println("Nuevo item añadido a la lista:", message)

		pageVariables := PageVariables{
			Title:   "Checklist por Gabriel Alvarez",
			Message: "Nuevo item añadido a la lista: " + message,
		}

		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, pageVariables)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}