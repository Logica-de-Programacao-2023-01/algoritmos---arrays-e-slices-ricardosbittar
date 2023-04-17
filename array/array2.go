package main

import "fmt"

func main() {

	// Criando um array de floats com 6 elementos
	numeros := [6]float64{3.14, 2.71, 1.618, 0.577, 1.414, 2.236}

	var valor float64
	fmt.Print("Informe um número: ")
	fmt.Scan(&valor)

	// Adicionando o valor em todas as posições do array
	for i := 0; i < len(numeros); i++ {
		numeros[i] += valor
	}

	// Imprimindo o array resultante
	fmt.Println("%f.2", numeros)
}
