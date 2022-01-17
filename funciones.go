package main

// Leer contenido de archivo .txt
import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ups, el programa no finalizo de manera correcta!")
		}
	}()

	if file, err := os.Open("example.txt"); err != nil {
		panic("no fue posible obtener el archivo!!")
	} else {
		defer func() {
			fmt.Println("Cerramos el archivo!!")
			file.Close()
		}()

		contenido := make([]byte, 254)

		longitud, _ := file.Read(contenido)

		contenidoArchivo := string(contenido[0:longitud])

		fmt.Println(contenidoArchivo)
	}

}

/*
// programar funciones

func funcion1() {
	fmt.Println(("Hola, soy la primera funcion!!!"))
}

func funcion2() {
	fmt.Println(("Hola, soy la segunda funcion!!!"))
}

func main() {

	fmt.Println("Hola nos encontramos en la funcion Main!!!")

	defer funcion1() // manda la funcion a que se ejecute de ultimas asi este de primeras o segundas ...

	funcion2()
}
*/

/*
//Variables globales
var username string // se inicializa como ""

func funcion1() {
	fmt.Println("Funcion1:", username)
}

func funcion2() {
	fmt.Println("Funcion2:", username)
}

func main() {
	username = "Cody"
	funcion1()
	funcion2()
}
*/

/*
// Manejo de errores
import "fmt"
import "errors"

func division(dividendo, divisor int) (int, error) {

	if divisor == 0 {
		return 0, errors.New("No se puede dividir por 0, ingrese otro numero")
	} else {
		return dividendo / divisor, nil
	}

}

func main() {

	if resultado, err := division(10, 0); err != nil {
		fmt.Println((err))
	} else {
		fmt.Println("El resultado es:", resultado)
	}
}
*/

/*
// Funciones como valores

func factorial(numero int) int {
	if numero == 1 {
		return 1
	}

	return factorial(numero-1) * numero
}

type customFunction func(n int) int

func main() {
	//var miFuncion = factorial
	var miFuncion customFunction

	if miFuncion == nil {
		miFuncion = factorial
	}

	resultado := miFuncion(3)

	fmt.Println(resultado)

}
*/

/*

// Recursion

func factorial(numero int) int {
	if numero == 1 {
		return 1
	}

	return factorial(numero-1) * numero
}

func main() {
	resultado := factorial(5)

	fmt.Println("El factorial es:", resultado)
}

*/

/*
// punteros

func modifiarValor(parametro *string) {
	*parametro = "Cambio de valor"
}

func main() {
	var curso = "Curso profesional de Go!"

	fmt.Println("Antes de la funcion:", curso)

	modifiarValor(&curso) // referencia

	fmt.Println("Despues de la funcion:", curso)
}
*/

/*
// Variadic functions o funciones variadicas

func promedio(calificaciones ...int) int { // los ... sirve para que la funcion reciba n cantidad de parametros

	var promedio, suma int

	for_, calificacion := range calificaciones {
		suma = suma + calificacion
	}
	promedio = suma / len(calificaciones)

	return promedio
}
func main() {

	resultado := promedio(10, 4, 1, 10)
	// variadic function
	fmt.Println("Promedio:", resultado)
}
*/
/*
// varibales y scopes

func modificarVariable(parametro string) {
	fmt.Println("Valor actual:", parametro)
	parametro = "Cambio de valor"
	fmt.Println("Nuevo Valor:", parametro)

	fmt.Printf("%p\n", &parametro)
}

func main() {

	var curso = "Curso profesional de Go!"

	modificarVariable(curso)

	fmt.Println(">>>", curso)
}
*/

/*
// bloques y alcances
// un bloque es una coleccion de sentencias

func main() {
	// bloque 1

	var curso = "Curso profesional de Go!"
	var version = 15
	{
		// bloque 2
		fmt.Println(version)
		fmt.Println(curso)
		{
			// bloque 3
			var version = 5
			fmt.Println(version)
			fmt.Println(curso)
		}
	}
	fmt.Println(curso)
}
*/

/*
// retornar funciones con otras funciones

type Operacion func(balance, cantidad int) int

func crearOperacion(tipo string) Operacion {
	if tipo == "suma" {
		return func(balance, cantidad int) int {
			return balance + cantidad
		}
	} else if tipo == "resta" {
		return func(balance, cantidad int) int {
			return balance - cantidad
		}
	} else {
		return func(balance, cantidad int) int { return balance * cantidad }
	}
}

func main() {
	suma := crearOperacion("suma")

	resultado := suma(40, 50)

	fmt.Println("el resultado es: ", resultado)
}
*/

/*
// funciones como argumentos
// se crea un nuevo tipo de dato con type

type Operacion func(balance, cantidad int) int // ejemplo debe retornar 2 parametros

func suma(num1, num2 int) int { // pasa a ser una funcion de tipo operacion
	return num1 + num2
}

func resta(num1, num2 int) int { // pasa a ser una funcion de tipo operacion
	return num1 - num2
}

func ejecutarOperacion(funcion Operacion, balance, cantidad int) {
	fmt.Println("Antes de la operacion")

	resultado := funcion(balance, cantidad)
	fmt.Println("El resultado es:", resultado)

	fmt.Println("Despues de la operacion")
}

func main() {
	ejecutarOperacion(suma, 100, 50)
}
*/

/*
// funciones anonimas

func main() {

	func() {
		fmt.Println("Hola mundo desde una func anonima! ")
	}() // se agregan los parentesis al final para no generar error

	miFuncion := func(username string) {
		fmt.Println("Hola", username, "desde una func anonima! ")
	}
	miFuncion("Cody")

	miFuncion2 := func(username string) string {
		message := fmt.Sprintf("Hola %s desde una func anonima! ", username)
		return message
	}

	mensajeFInal := miFuncion2("Gopher")

	fmt.Println(mensajeFInal)
}
*/
/*
// retornar valores apartir de funciones

func suma(num1, num2 int) (int, string) { // el segundo int indica que es un tipo de dato en comun para las dos variables y el ultimo int es el tipo de dato a retornar
	return num1 + num2, "funcion suma"
}

func main() {
	resultado, mensaje := suma(4, 2) // o tambien se puede usar _ resultado, _ := suma(4, 2)

	fmt.Println(resultado, mensaje)
}
*/

/*
// funcion que saluda a los usuarios
func saluda(username string) {
	fmt.Println("Hola", username)
}

func main() {
	saluda("Cody")
}
*/
