package main

import "fmt"

type UserType int

const (
	Teacher UserType = 1
	Student UserType = 2
)

type User struct {
	Username string
	Type     UserType
}

func main() {

	steven := User{Username: "steven", Type: Student}
	alvaro := User{Username: "Alvaro", Type: Teacher}

	fmt.Println(steven)
	fmt.Println(alvaro)

	if steven.Type == Student {
		fmt.Println("El usuario", steven.Username, "es un estudiante")
	}
}
