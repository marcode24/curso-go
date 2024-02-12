package main

import "fmt"

func main() {
	// slice no poseen daots, son apuntadores a un array
	set := [7]string{"🐶", "🐱", "🐭", "🐹", "🐰", "🦊", "🐻"}
	animals := set[0:5]
	others := set[2:7]

	// si no se especifica el inicio o el final, se toma el inicio o el final del array
	// animals := set[:5]
	// others := set[2:]

	// si no se especifica ambos, se toma el array completo
	// animals := set[:]
	// others := set[:]

	fmt.Println(animals)
	fmt.Println(others)

	// len es el tamano del slice
	fmt.Println("Tamano de animals:", len(animals))

	// capacidad es el tamano del array desde el inicio del slice
	fmt.Println("Capacidad de animals:", cap(animals))

	// append agrega un elemento al slice
	animals = append(animals, "🐻")
	fmt.Println(animals)
	fmt.Println("Tamano de animals:", len(animals))
	fmt.Println("Capacidad de animals:", cap(animals))

	// si se excede la capacidad, se crea un nuevo array con el doble de capacidad
	animals = append(animals, "🐼")
	fmt.Println(animals)
	fmt.Println("Tamano de animals:", len(animals))
	fmt.Println("Capacidad de animals:", cap(animals))

	// fruits := []string{"🍎", "🍌", "🍇", "🍓", "🍑"}
	fruits := make([]string, 0, 5)

	// fruits = append(fruits, "🍍")
	fmt.Println(fruits)
	fmt.Println("Tamano de fruits:", len(fruits))
	fmt.Println("Capacidad de fruits:", cap(fruits))

}
