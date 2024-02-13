package main

import (
	"fmt"
	"os"
)

func main() {
	// defer fmt.Println("world")

	// fmt.Println("hello")
	// salida: hello world

	// a := 5
	// defer fmt.Println("defer 1:", a)
	// a = 10
	// fmt.Println("defer 2:", a)

	// salida:
	// defer 2: 10
	// defer 1: 5

	// uso practico
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo", err)
		return
	}
	// Se ejecuta al final de la funcion
	// Esto nos asegura que el archivo se cierre al finalizar la funcion
	defer file.Close()

	_, err = file.Write([]byte("Hola mundo"))
	if err != nil {
		fmt.Println("Error al escribir en el archivo", err)
		return
	}

	division(10, 2)
	division(10, 0)
	division(10, 3)
}

func division(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado en division", r)
		}
	}()
	validarDivisor(divisor)
	fmt.Println(dividendo / divisor)
}

func validarDivisor(divisor int) {
	if divisor == 0 {
		panic("No se puede dividir por cero")
	}
}
