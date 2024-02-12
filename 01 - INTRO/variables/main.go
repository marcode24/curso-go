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

	// Uso de comentarios
	// Este es un comentario de una sola línea
	/*
		Este es un comentario de varias líneas
		Este puede ser útil para explicar el código
	*/

	// tipos de datos
	// booleanos
	var a bool = true
	fmt.Printf("Tipo: %T, Valor: %v \n", a, a)

	// string
	var s string = "Hola Mundo"
	fmt.Printf("Tipo: %T, Valor: %v \n", s, s)

	// numericos
	var edad uint8 = 255
	fmt.Printf("Tipo: %T, Valor: %v \n", edad, edad)

	/*
		NOTA:
		'byte' es un alias para uint8
		'rune' es un alias para int32, también representa un punto de código Unicode
	*/

	var edad2 byte = 255
	fmt.Printf("Tipo: %T, Valor: %v \n", edad2, edad2)

	var edad3 rune = 2147483647
	fmt.Printf("Tipo: %T, Valor: %v \n", edad3, edad3)

	var uni rune = 'a'
	fmt.Printf("Tipo: %T, Valor: %v \n", uni, uni)

	var flotante float32 = 3.14
	fmt.Printf("Tipo: %T, Valor: %v \n", flotante, flotante)

	var flotante2 float64 = 3.1416
	fmt.Printf("Tipo: %T, Valor: %v \n", flotante2, flotante2)

	/*
		- uint8: 0 a 255
		- uint16: 0 a 65535
		- uint32: 0 a 4294967295
		- uint64: 0 a 18446744073709551615
		- int8: -128 a 127
		- int16: -32768 a 32767
		- int32: -2147483648 a 2147483647
		- int64: -9223372036854775808 a 9223372036854775807
		- float32: IEEE-754 32-bit floating-point numbers
		- float64: IEEE-754 64-bit floating-point numbers
	*/

	var p uint8 = 100
	var o uint16 = 23000
	c := uint16(p) + o
	fmt.Printf("Tipo: %T, Valor: %v \n", c, c)

	// operador blank
	// El operador blank (_) es una variable que se utiliza para descartar valores
	// que no necesitamos
	_ = "🐶"
}
