package main

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

// Declaracion de alias
type myAlias = course

// Definicion de un tipo de dato
type newCourse course

type newBool bool

func (b newBool) String() string {
	if b {
		return "VERDADERO"
	}
	return "FALSO"
}

func main() {
	// c := myAlias{name: "Go"}
	c := newCourse{name: "Go"}
	// c.Print()
	fmt.Printf("El tipo de dato de c es: %T\n", c)

	var b newBool = true
	fmt.Println(b.String())
}
