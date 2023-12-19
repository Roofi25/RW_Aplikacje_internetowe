package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func stronaZad3Func(w http.ResponseWriter, r *http.Request) {
	/*
		randomNumber := rand.Intn(3)
		if randomNumber == 0 {
			http.ServeFile(w, r, "strona1.html")
		} else if randomNumber == 1 {
			http.ServeFile(w, r, "strona2.html")
		} else if randomNumber == 2 {
			http.ServeFile(w, r, "strona3.html")
		}
	*/

	randomNumber := rand.Intn(3) + 1
	randomNumberToString := fmt.Sprint(randomNumber)
	http.ServeFile(w, r, "strona"+randomNumberToString+".html")
}

func main() {
	http.HandleFunc("/zad3/", stronaZad3Func)
	http.ListenAndServe(":10000", nil)
}
