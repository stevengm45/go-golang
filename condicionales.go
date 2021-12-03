package main

import (
	"fmt"
)

func main() {

	// Obtener valores de un mapa

	usuarios := make(map[string]string)

	usuarios["steven"] = "stevengm45@gmail.com"

	if correo, ok := usuarios["cody"]; ok { // por convencion se nombre la variable ok la cual almacena true o false
		fmt.Println(correo)
	} else {
		fmt.Println("No se pudo encontrar el valor")
	}

	/*
		// switch
		// 10 -9/8 - 7/6 - 5/0

		var nota int
		fmt.Print("ingrese la nota del estudiante: ")
		fmt.Scanf("%d", &nota)

		switch nota {
		case 10:
			fmt.Println("ACalificacion perfecta, eres un genio!")
		case 8, 9:
			fmt.Println("Aprobaste la materia!")
		case 6, 7:
			fmt.Println("Aprobaste pero muy mediocre!")
		case 5, 4, 3, 2, 1, 0:
			fmt.Println("Reprobaste la materia!")
		default:
			fmt.Println("Calificación no valida")
		}
	*/
	/*
		switch {
		case nota == 10:
			fmt.Println("Felicitaciones calificacion perfecta!")
		case nota == 9 || nota == 8:
			fmt.Println("Aprobaste la materia!")
		case nota == 7 || nota == 6:
			fmt.Println("Aprobaste pero muy mediocre!")
		case nota <= 5 && nota >= 0:
			fmt.Println("Reprobaste la materia!")
		default:
			fmt.Println("Calificación no valida")
		}
	*/

	/*
		// Declaracion de variables en condiciones

		if nombre, edad := "Cody", 7; nombre == "Cody " {
			fmt.Println("Hola", nombre, "te damos la bienvenida al curso de Go!")
		} else {
			fmt.Println("los valores son: ", nombre, edad)
		}
	*/

	/*

		//else if
		// 10 - 8/9 - 6/7 - 5

		var nota float32
		fmt.Print("ingrese la nota del estudiante: ")
		fmt.Scanf("%f", &nota)

		if nota == 10 {
			fmt.Println("La calificacion es perfecta")
		} else if nota <= 9 && nota >= 8 {
			fmt.Println("Aprobaste la materia")
		} else if nota < 8 && nota >= 6 {
			fmt.Println("la calificacion es mediocre pero paso")
		} else if nota < 6 && nota >= 0 {
			fmt.Println("reprobaste")
		} else {
			fmt.Println("Calificacion no valida")
		}
	*/
	/*
		var edad int
		fmt.Print("ingresa tu edad: ")
		fmt.Scanf("%d", &edad)

		if edad >= 18 {
			fmt.Println("Es mayor de edad")
		} else {
			fmt.Println("Es menor de edad")
		}
	*/
}
