package main

import "fmt"

func main() {

	//Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro. Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante
	var valor int

	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(numeros)

	fmt.Println("Informe um valor")

	fmt.Scanln(&valor)

	presente := false

	for i := 0; i < len(numeros); i++ {

		if numeros[i] == valor {

			presente = true

		}
	}

	if presente {
		fmt.Println("O valor que voce escolheu ja estava dentro do conjunto antes")
	} else {
		novoslice := append(numeros, valor)
		fmt.Println("O novo valor foi adicionado ao conjunto, o qual ficou o seguinte:", novoslice)
	}

}
