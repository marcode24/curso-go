package main

import "fmt"

func main() {
	var food [3]string
	food[0] = "🍔"
	food[1] = "🍟"
	food[2] = "🍦"
	// food[3] = "🍕" // Esto generaría un error de índice fuera de rango

	fmt.Println(food)

	foods := [3]string{"🍔", "🍟", "🍦"}

	fmt.Println(foods)
}
