package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/patrikeh/go-deep"
	"github.com/patrikeh/go-deep/training"
)

var n *deep.Neural
var trainingData *training.Examples

func uczenieFunc(w http.ResponseWriter, r *http.Request) {

	rand.Seed(time.Now().UTC().UnixNano())

	vars := mux.Vars(r)
	input1, _ := strconv.ParseFloat(vars["input1"], 64)
	input2, _ := strconv.ParseFloat(vars["input2"], 64)

	// dane uczące
	var data = training.Examples{
		{Input: []float64{input1, input2}, Response: []float64{0}},
	}

	trainingData = &data

	// optymalizacja sieci neuronowej
	optimizer := training.NewSGD(0.05, 0.1, 1e-6, true)
	trainer := training.NewTrainer(optimizer, 50)
	training, heldout := data.Split(0.5)

	//zmienna globalna n zeby jej nie redeklarować
	trainer.Train(n, training, heldout, 1000)

	w.Header().Set("Content-Type", "application/json") // nagłówek JSON

	// wyświetlenie predykcji dla danych
	blad := 0.0
	for i := 0; i < len(data); i++ {
		predykcja := n.Predict(data[i].Input)
		prawidlowo := data[i].Response
		blad += math.Abs(predykcja[0] - prawidlowo[0])
		data, _ := json.Marshal(predykcja) // konwersja na JSON
		w.Write(data)                      // zwrócenie danych JSON
	}
}

func resetujFunc(w http.ResponseWriter, r *http.Request) {
	//uczymy sieć na nowo
	n = deep.NewNeural(&deep.Config{
		Inputs:     2,
		Layout:     []int{2, 2, 1},
		Activation: deep.ActivationSigmoid,
		Mode:       deep.ModeBinary,
		Weight:     deep.NewNormal(1.0, 0.0),
		Bias:       true,
	})

	errorAfterTraining := liczBlad(n, *trainingData)
	fmt.Fprintf(w, "<html><body>Sieć neuronowa została zresetowana! Błąd sieci: %v</body></html>", errorAfterTraining)
}

func liczBlad(net *deep.Neural, data training.Examples) float64 {
	var blad float64
	for _, example := range data {
		prediction := net.Predict(example.Input)
		blad += math.Abs(prediction[0] - example.Response[0])
	}
	return blad / float64(len(data))
}

func main() {

	// utworzenie sieci neuronowej (2 wejścia, 2-2-1 neuronów)
	//zmienna globalna (uczona tylko raz chyba ze będzie /resetuj/)
	n = deep.NewNeural(&deep.Config{
		Inputs:     2,
		Layout:     []int{2, 2, 1},
		Activation: deep.ActivationSigmoid,
		Mode:       deep.ModeBinary,
		Weight:     deep.NewNormal(1.0, 0.0),
		Bias:       true,
	})

	r := mux.NewRouter()
	//   http://localhost:10000/predykcja/2.7810836/2.550537003
	r.HandleFunc("/predykcja/{input1}/{input2}", uczenieFunc)
	//   http://localhost:10000/resetuj
	r.HandleFunc("/resetuj", resetujFunc)
	http.ListenAndServe(":10000", r)
}

/*
// inicjalizacja ziarna losowania
	rand.Seed(time.Now().UTC().UnixNano())
	// dane uczące
	var data = training.Examples{
		{Input: []float64{2.7810836, 2.550537003}, Response: []float64{0}},
		{Input: []float64{1.465489372, 2.362125076}, Response: []float64{0}},
		{Input: []float64{3.396561688, 4.400293529}, Response: []float64{0}},
		{Input: []float64{1.38807019, 1.850220317}, Response: []float64{0}},
		{Input: []float64{7.627531214, 2.759262235}, Response: []float64{1}},
		{Input: []float64{5.332441248, 2.088626775}, Response: []float64{1}},
		{Input: []float64{6.922596716, 1.77106367}, Response: []float64{1}},
		{Input: []float64{8.675418651, -0.242068655}, Response: []float64{1}},
	}
	// utworzenie sieci neuronowej (2 wejścia, 2-2-1 neuronów)
	n := deep.NewNeural(&deep.Config{
		Inputs:     2,
		Layout:     []int{2, 2, 1},
		Activation: deep.ActivationSigmoid,
		Mode:       deep.ModeBinary,
		Weight:     deep.NewNormal(1.0, 0.0),
		Bias:       true,
	})
	// optymalizacja sieci neuronowej
	optimizer := training.NewSGD(0.05, 0.1, 1e-6, true)
	trainer := training.NewTrainer(optimizer, 50)
	training, heldout := data.Split(0.5)
	trainer.Train(n, training, heldout, 1000)
	// wyświetlenie predykcji dla danych
	blad := 0.0
	for i := 0; i < len(data); i++ {
		predykcja := n.Predict(data[i].Input)
		prawidlowo := data[i].Response
		blad += math.Abs(predykcja[0] - prawidlowo[0])
		fmt.Println("PREDYKCJA:", predykcja, ", PRAWIDŁOWO:", prawidlowo)
		fmt.Println(", BLAD:", blad)
	}
	fmt.Println("BŁĄD (SUMA):", blad)
*/
