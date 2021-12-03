package main

import "fmt"

func main() {

	var numero = 1
	for {
		fmt.Println(numero)
		numero++

		if numero == 100 {
			break // tambien se puede usar panic, ejm = panic("fin del ciclo")
		}
	}

	/*
		// do while
		var numero = 10

		for ok := true; ok; ok = numero <= 10 { // do while
			fmt.Println(numero)
			numero++
		}
	*/
	/*
		//break - continue

		for i := 1; i <= 10; i++ {

			if i == 5 {
				continue
			}

			if i == 8 {
				break
			}

			fmt.Println((i))
		}
	*/
	/*
		// foreach

		animales := [...]string{"Perro", "Gato", "Pez", "Vaca", "Cerdo"}

		for indice, elemento := range animales {
			fmt.Println("El indice es:", indice, "El valor es:", elemento)
		}

		for _, elemento := range animales {
			fmt.Println("El valor es:", elemento)
		}

		for indice := 0; indice < len(animales); indice++ {
			elemento := animales[indice]
			fmt.Print(elemento)
		}
	*/
	/*
		//for como while
		numero := 12345
		contador := 0

		for numero > 0 {
			numero /= 10
			contador++
		}

		fmt.Println("La cantidad de digitos es: ", contador)
	*/
	/*
		// for variable; condicion; incremento
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	*/
}
