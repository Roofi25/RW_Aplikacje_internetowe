package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func loadCars() []*car {
	cars := []*car{}
	file, err := os.Open("cars.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		c := car{}
		c.mpg, _ = strconv.ParseFloat(line[0], 64)
		c.cylinders, _ = strconv.Atoi(line[1])
		c.displacement, _ = strconv.ParseFloat(line[2], 64)
		c.horsepower, _ = strconv.ParseFloat(line[3], 64)
		c.weight, _ = strconv.ParseFloat(line[4], 64)
		c.acceleration, _ = strconv.ParseFloat(line[5], 64)
		c.year, _ = strconv.Atoi(line[6])
		c.origin, _ = strconv.Atoi(line[7])
		c.name = line[8]
		cars = append(cars, &c)
	}
	return cars
}

// funkcja porównująca obiekty przekazane przez parametry
func compare(first *car, second *car) float64 {
	if first == second {
		return 0.0
	}
	similarity := 1.0
	similarity *= 1.0 - math.Abs(first.mpg-second.mpg)/40
	similarity *= 1.0 - math.Abs(first.horsepower-second.horsepower)/300
	similarity *= 1.0 - math.Abs(first.weight-second.weight)/5000
	similarity *= 1.0 - math.Abs(first.acceleration-second.acceleration)/30
	return similarity
}

// funkcja jak metoda - porównuje obiekt bieżący z podanym jako parametr
func (this *car) compare(other *car) float64 {
	if this == other {
		return 0.0
	}
	similarity := 1.0
	similarity *= 1.0 - math.Abs(this.mpg-other.mpg)/40
	similarity *= 1.0 - math.Abs(this.horsepower-other.horsepower)/300
	similarity *= 1.0 - math.Abs(this.weight-other.weight)/5000
	similarity *= 1.0 - math.Abs(this.acceleration-other.acceleration)/30
	return similarity
}

type car struct {
	mpg          float64
	cylinders    int
	displacement float64
	horsepower   float64
	weight       float64
	acceleration float64
	year         int
	origin       int
	name         string
}

func main() {
	car0 := car{18, 8, 307, 130, 3504, 12, 70, 1, "chevrolet malibu"}
	car1 := car{13, 8, 351, 158, 4363, 13, 73, 1, "ford ltd"}
	car2 := car{29, 4, 98, 83, 2219, 16.5, 74, 2, "audi fox"}
	car3 := car{20, 6, 232, 100, 2914, 16, 75, 1, "amc gremlin"}
	car4 := car{33, 4, 91, 53, 1795, 17.4, 76, 3, "honda civic"}
	car5 := car{23.2, 4, 156, 105, 2745, 16.7, 78, 1, "plymouth sapporo"}
	cars := []car{car0, car1, car2, car3, car4, car5}
	fmt.Println("zad 2-3")
	fmt.Println(cars)

	fmt.Println("zad 4 (test nie na wskaznikach)")
	cars[4].name = "test"
	fmt.Println(cars[4])
	fmt.Println(car4)

	//zad5
	cars_pointers := []*car{&car0, &car1, &car2, &car3, &car4, &car5}
	fmt.Println("zad 6 (test na wskaźnikach)")
	cars_pointers[4].name = "test"
	fmt.Println(cars_pointers[4])
	fmt.Println(car4)

	//zad8
	fmt.Println("zad 8")
	wynik1 := compare(&car0, &car1)
	wynik2 := car0.compare(&car1)
	fmt.Println(wynik1, wynik2)

	//zad9
	maks_value := cars[0].compare(&car2)
	maks_car := cars[0]
	for i := 1; i < len(cars); i++ {
		if cars[i] != car2 {
			result := cars[i].compare(&car2)
			if result > maks_value {
				maks_car = cars[i]
				maks_value = result
			}
		}
	}
	fmt.Println("zad 9 (najbardziej podobny samochod do car2)")
	fmt.Println("Samochód: ", maks_car, " podobieństwo: ", maks_value)

	//zad10
	cars_z_pliku := loadCars()
	maks_value_zad10 := cars_z_pliku[0].compare(cars_z_pliku[2])
	maks_car_zad10 := cars_z_pliku[0]
	for i := 1; i < len(cars); i++ {
		if cars_z_pliku[i] != cars_z_pliku[2] {
			result := cars_z_pliku[i].compare(cars_z_pliku[2])
			if result > maks_value {
				maks_car_zad10 = cars_z_pliku[i]
				maks_value_zad10 = result
			}
		}
	}
	fmt.Println("zad 10 (najbardziej podobny samochod do trzeciego z wczytanych)")
	fmt.Println("Samochód: ", maks_car_zad10, " podobieństwo: ", maks_value_zad10)
}
