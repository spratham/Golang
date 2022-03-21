package main

import "fmt"

func main() {

	Pratham := User{"Pratham", "pratham919singh@gmail.com", true, 16}
	fmt.Printf("Pratham Details- %+v\n", Pratham)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
