package main

import "fmt"

func main() {
	/*
			var numeros [5]int

			numeros[0] = 100

			fmt.Print(numeros, "\n")

		numeros := [5]int{100, 200, 300, 400, 500}

		fmt.Println(numeros)

		numeros2 := [...]int{100, 200, 500}

		fmt.Println(numeros2)
	*/

	monedas := [...]string{3: "Peso Colombiano", 2: "Dolar", 1: "Euro", 0: "Libras"}

	fmt.Println(monedas[0])
	fmt.Println(monedas[1])
	fmt.Println(monedas[2])
	fmt.Println(monedas[3])

	subMonedas := monedas[0:3]
	fmt.Println(subMonedas)
}
