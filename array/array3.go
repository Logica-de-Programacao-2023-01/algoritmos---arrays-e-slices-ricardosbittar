package main

import "fmt"

func main() {

	//Crie um Array de floats com 10 elementos. Crie um novo Slice que contenha apenas os elementos do Array que são maiores que 5. Imprima o novo Slice

	array := [10]float64{1.5, 4.7, 8.7, 9.5, 9.4, 10.6, 4.6, 3.3, 2.1, 1.4}
	slice := []float64{}

	for _, i := range array {

		if i > 5 {
			slice = append(slice, i)

		}
	}

	fmt.Println(slice)
}
