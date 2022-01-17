package main

import "fmt"

type Animal interface {
	Comer()
	Dormir()
}

type Perro struct {
	Nombre string
}

func (self Perro) Comer() {
	fmt.Println("El perro come!")
}

func (self Perro) Dormir() {
	fmt.Println("El perro duerme!")
}

func ejecutarAcciones(animal Animal) {
	animal.Comer()
	animal.Dormir()
}

func main() {
	totoy := Perro{Nombre: "Toto"}

	ejecutarAcciones(totoy)
}
