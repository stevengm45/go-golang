package main

import "fmt"

type Curso struct {
	Titulo string
	Videos []Video
}

type Video struct {
	Titulo string
	Curso  Curso
}

func main() {

	video1 := Video{Titulo: "1-Introduccion"}
	video2 := Video{Titulo: "2-Instalacion"}

	videos := []Video{video1, video2}

	curso := Curso{Titulo: "Curso profesional de Go!", Videos: videos}

	video1.Curso = curso
	video2.Curso = curso

	fmt.Println(video1.Curso.Titulo)

	for _, video := range curso.Videos { // iterar el arreglo
		fmt.Println(video.Titulo)
	}

}

/*
// Relacion uno a uno
// Estructuras embebidas

type User struct {
	Name   string
	Email  string
	Active bool
}

type Student struct {
	User User
	Id   string
}

func main() {

	steven := User{Name: "Steven", Email: "steven@codigo.com", Active: true}

	alvaro := User{Name: "Alvaro", Email: "pipe@codigo.com", Active: true}

	stevenStudent := Student{User: steven, Id: "1f1"}

	fmt.Print(alvaro)
	fmt.Println(stevenStudent.User.Active)
}
*/
/*

// Crear Metodos

type User struct {
	Name  string
	Email string
}

func (self *User) setName(name string) { // de esta manera se convierte la funcion en un metodo
	self.Name = name
}

func (self *User) getName() string {
	return self.Name
}

func main() {

	cody := User{Name: "Cody", Email: "info@codigo.com"}

	cody.setName("CodigoFacilito")

	fmt.Println(cody.getName())

}
*/

/*

// Crear estructuras

type User struct {
	Name  string
	Email string
}

func main() {
	var cody User // Un objeto

	cody.Name = "Cody"
	cody.Email = "info@codigofacilito.com"

	cody2 := User{Name: "Cody2", Email: "cody@@codigo.com"}

	Name := "Cody3"
	Email := "cody3@codigo.com"

	cody3 := User{Name, Email}

	fmt.Println(cody)
	fmt.Println(cody.Name)
	fmt.Println(cody.Email)

	fmt.Println(cody2)
	fmt.Println(cody2.Name)
	fmt.Println(cody2.Email)

	fmt.Println(cody3)
	fmt.Println(cody3.Name)
	fmt.Println(cody3.Email)
}
*/
