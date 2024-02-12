package main

import "fmt"

func main() {
	animals := make(map[string]string)
	animals["🐶"] = "Perro"
	animals["🐱"] = "Gato"
	fmt.Println(animals)

	fruits := map[string]string{
		"apple":  "🍎",
		"banana": "🍌",
	}
	fmt.Println(fruits)

	// eliminar un elemento
	delete(fruits, "apple")
	fmt.Println(fruits)

	// obtener un elemento
	apple, ok := fruits["banana"]
	fmt.Println(apple, ok)

	// obtener un elemento que no existe
	orange, ok := fruits["orange"]
	fmt.Println(orange, ok)

	// si no existe, se agrega el elemento
	if _, ok := animals["lion"]; !ok {
		animals["lion"] = "León"
	}

	fmt.Println(animals)
}
