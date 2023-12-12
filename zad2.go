package main

import "fmt"

func minmax(tab []int) (int, int) {
	if len(tab) < 1 {
		return 0, 0
	}
	min, max := tab[0], tab[0]
	// w tym miejscu należy napisać pętlę
	// wyszukującą minimum i maksimum
	// z tablicy tab
	for i := 1; i < len(tab); i++ {
		if tab[i] < min {
			min = tab[i]
		} else if tab[i] > max {
			max = tab[i]
		}
	}
	return min, max
}
func main() {
	numbers := []int{10, 2, 24, 13, 20}
	a, b := minmax(numbers)
	fmt.Println("Min: ", a, "Max: ", b)
}
