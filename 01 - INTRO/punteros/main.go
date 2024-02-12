package main

import "fmt"

func main() {
	fruit := "🍎"
	var pointer *string = &fruit
	// Para obtener la dirección de memoria de una variable, usamos el operador '&'
	fmt.Printf("Tipo: %T, Valor: %v, Dirección: %v \n", fruit, fruit, &fruit)
	// Para obtener el valor almacenado en una dirección de memoria, usamos el operador '*'
	fmt.Printf("Tipo: %T, Valor: %v, Dirección: %v \n", pointer, *pointer, pointer)
	// Obtener el valor con desreferencia
	fmt.Printf("Tipo: %T, Valor: %v, Desreferencia: %v \n", pointer, pointer, *pointer)

	// Cambiando el valor de la variable original
	fruit = "🍌"
	fmt.Printf("Tipo: %T, Valor: %v, Dirección: %v \n", fruit, fruit, &fruit)
	// El valor de la variable original cambió, pero el valor almacenado en la dirección de memoria no cambió
	fmt.Printf("Tipo: %T, Valor: %v, Dirección: %v \n", pointer, *pointer, pointer)
	// Obtener el valor con desreferencia
	fmt.Printf("Tipo: %T, Valor: %v, Desreferencia: %v \n", pointer, pointer, *pointer)
}
