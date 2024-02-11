package main

import "fmt"

func main() {

	// Declaración de variables con su tipo
	var dogName string = "Max"

	// Creando una variable y asignando un valor
	var dog string
	dog = "🐶"

	fmt.Println(dogName)
	fmt.Println(dog)

	// Declaracion de varias variables
	var doggy, cat string = "🐶", "🐱"
	fmt.Println(doggy, cat)

	// Declaracion de vaiables sin tipo de dato
	var dog2 = "🐶"
	fmt.Println(dog2)

	// Declaracion de variables con shorthand
	dog3 := "🐶"
	fmt.Println(dog3)

	// Uso de variables constantes
	const PI = 3.14
	// Si intentamos cambiar el valor de una constante, Go nos arrojará un error
	// PI = 3.1416
	fmt.Println(PI)

}
