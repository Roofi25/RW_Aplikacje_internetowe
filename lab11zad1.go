package main

import (
	"net/http"
	"text/template"
)

type Student struct {
	Imie     string
	Nazwisko string
	Indeks   int
	Email    string
}

func stronaFunc(w http.ResponseWriter, r *http.Request) {
	// zwrócenie statycznej strony strona.html
	//http.ServeFile(w, r, "dane.html")
	tmpl, _ := template.ParseFiles("dane.html")
	tmpl.Execute(w, studenci)
}

var studenci []Student = []Student{
	{"Jan", "Kowalski", 12345, "test@test"},
	{"Marek", "Nowak", 30000, "to@tamto"},
	{"Anna", "Zdyb", 23232, "anna@zdyb"},
}

func main() {
	http.HandleFunc("/test/", stronaFunc)
	http.ListenAndServe(":10000", nil)
}
