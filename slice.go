package main

import "fmt"

func main() {

	slice := make([]int, 3, 3)

	slice[0] = 100
	slice[1] = 200
	slice[2] = 300

	slice = append(slice, 400)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	/*
		numeros := []int{1, 2, 3, 4, 5, 6}

		inicio := numeros[0:3]
		final := numeros[3:6]

		numeros[0] = 100
		numeros[5] = 600

		fmt.Println(numeros)
		fmt.Println(inicio)
		fmt.Println(final)
	*/

	/*
		meses := []string{"enero", "febrero", "marzo", "abril", "mayo", "junio"}

		//puntero
		//longitud
		// una capacidad

		longitud := len(meses)
		capacidad := cap(meses)

		fmt.Printf("La longitud es: %v, La capacidad es: %v %p\n", longitud, capacidad, meses)

		meses = append(meses, "Julio", "Agosto") // si la estructura se encuentra en su capacidad maxima se genera un nuevo arreglo

		longitud = len(meses)
		capacidad = cap(meses)

		fmt.Println(meses)
		fmt.Printf("La longitud es: %v, La capacidad es: %v %p\n", longitud, capacidad, meses)
	*/

	/*
		numeros := []int{1, 2, 3, 4, 5}

		numeros = append(numeros, 6)

		nuevoSlice := numeros[0:5]

		numeros[0] = 100

		fmt.Println(numeros)
		fmt.Println(nuevoSlice)
	*/

}
