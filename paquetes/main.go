package main

import (
	"fmt"

	"./CodigoFacilito"
)

func main() {

	curso := CodigoFacilito.Curso{Titulo: "Curso profesional de Go"}

	fmt.Println(curso.ObtenerTitulo())
}
