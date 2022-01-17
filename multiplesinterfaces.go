package main

import "fmt"

type Animal interface {
	Dormir()
}

type Mascota interface {
	Comer()
}

type Gato struct {
	Nombre string
}

func (self Gato) Dormir() {
	fmt.Println("El gato duerme!")
}

func (self Gato) Comer() {
	fmt.Println("El gato come!")
}

func accionAnimal(animal Animal) {
	fmt.Println("El objeto es un animal")
}

func accionMascota(mascota Mascota) {
	fmt.Println("El objeto es una mascota")
}

func main() {

	gato := Gato{Nombre: "Negro!"}
	gato.Comer()
	gato.Dormir()

	accionAnimal(gato)
	accionMascota(gato)
}
