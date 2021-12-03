package main

import (
	"fmt"
)

// go build main.go -> archivo ejecutable
// go run main.go

// declarar constantes
/*
const Curso string = "Curso profesional de Go"

const (
	Domingo int = iota + 1 // valor por default 0, se puede añadir 1 para empezar desde 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)
*/
func main() {

	var nombre string
	var edad int
	var altura float32

	fmt.Print("ingresa tu nombre: ")
	fmt.Scanf("%s", &nombre)
	fmt.Print("ingresa tu edad: ")
	fmt.Scanf("%d", &edad)
	fmt.Print("ingresa tu altura: ")
	fmt.Scanf("%f", &altura)

	fmt.Printf("Hola %s con una edad %d y una altura de %.2f\n", nombre, edad, altura)

	/*
		// var curso string = "Curso profesional de Go!" // 1
		// var curso = "Curso profesional de Go!" // 2
		curso := "Curso profesional de Go!"

		longitud := len(curso) // len sirve para el alrgo del string

		fmt.Println(curso)
		fmt.Println(longitud)

		primerCaracter := curso[3] // Char -> uint8 | acceder a una posicion en especifico y lo imprime en codigo ASCII
		ultimoCaracter := curso[len(curso)-1] //obtener el ultimo caracter

		fmt.Println(primerCaracter)
		fmt.Println(reflect.TypeOf(primerCaracter)) // conocer el tipo de variable o tipo de string

		fmt.Printf("%c\n", primerCaracter)
		fmt.Printf("%c\n", ultimoCaracter)
	*/
	/*
		fmt.Println(Curso)
		fmt.Println(Domingo)
		fmt.Println(Lunes)
		fmt.Println(Martes)
		fmt.Println(Miercoles)
		fmt.Println(Jueves)
		fmt.Println(Viernes)
		fmt.Println(Sabado)
	*/

	/*
		operadores logicos
		&& and
		|| or
		! negar
	*/
	/*
		resultado := 20 == 20 && 30 == 30 && 20 > 40
		fmt.Println(resultado)

		resultado2 := 20 != 20 || 30 != 30 || 20 > 40
		fmt.Println(resultado2)

		resultado3 := 20 == 20 && 30 < 100 && (20 < 40 || 100 == 90)
		fmt.Println(resultado3)

		resultado4 := !true
		fmt.Println(resultado4)
	*/
	/*
		operadores relacionales
		>
		<
		>=
		<=
		==
		!=
	*/
	/*
		edad := 10
		resultado := edad > 5

		nombre := "Cody"
		resultado1 := nombre == "Cody"

		fmt.Println(resultado)
		fmt.Println(resultado1)
	*/
	/*
			var nombre string
			var apellido string
			var pais string

		// var nombre, apellido, pais = "Steven", "Gonzalez", "Colombia"
		nombre, apellido, pais := "Steven", "Gonzalez", "Colombia"
		edad, altura := 26, 1.83

		fmt.Println(nombre, apellido, pais, edad, altura)
	*/

	/*
		var nombre = "Steven Gonzalez"
		var edad int

		nombre := "Steven Gonzalez!"
		edad := 26

		var altura = 1.83

		fmt.Println("Hola", nombre, "tu edad es", edad, "y tu altura es", altura)
		fmt.Println(edad)
		fmt.Println(altura)
	*/
}
