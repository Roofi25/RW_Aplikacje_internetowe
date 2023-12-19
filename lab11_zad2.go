package main

import (
	"encoding/json"
	"net/http"
	"text/template"
)

type Student struct {
	Imie     string
	Nazwisko string
	Indeks   int
	Email    string
}

type rms struct {
	N1 string `json:"Imie"`     // ta właściwość zostanie zwrócona jako Imie
	N2 string `json:"Nazwisko"` // ta właściwość zostanie zwrócona jako Nazwisko
	N3 int    `json:"Indeks"`   // ta właściwość zostanie zwrócona jako Indeks
	N4 string `json:"-"`        // ta właściwość nie zostanie zwrócona
}

func studentsToJSONFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // nagłówek JSON

	var studenciBezEmailu []rms

	for _, student := range studenci {
		studentBezEmailu := rms{
			N1: student.Imie,
			N2: student.Nazwisko,
			N3: student.Indeks,
		}
		studenciBezEmailu = append(studenciBezEmailu, studentBezEmailu)
	}

	data, _ := json.Marshal(studenciBezEmailu) // konwersja na JSON
	w.Write(data)                              // zwrócenie danych JSON
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
	http.HandleFunc("/json/", studentsToJSONFunc)
	http.ListenAndServe(":10000", nil)
}
