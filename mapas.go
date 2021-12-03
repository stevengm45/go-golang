package main

import "fmt"

func main() {

	//Iterar sobre mapas

	usuarios := map[int]string{} // se deja a un lado la funcion make

	usuarios[1] = "Usuario 1"
	usuarios[2] = "Usuario 2"
	usuarios[3] = "Usuario 3"
	usuarios[4] = "Usuario 4"

	for llave, valor := range usuarios {
		fmt.Println(llave, valor)
	}

	/*
		usuarios := make(map[string][]int)

		usuarios["Steven"] = []int{10, 9, 9, 8, 2}

		fmt.Println(usuarios)
	*/

	/*
		dias := make(map[int]string)

		dias[0] = "domingo"
		dias[1] = "lunes"
		dias[2] = "martes"
		dias[3] = "miercoles"
		dias[4] = "jueves"
		dias[5] = "viernes"
		dias[6] = "sabado"

		delete(dias, 4)

		fmt.Println(dias)
	*/
}
