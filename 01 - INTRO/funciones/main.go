package main

import (
	"errors"
	"fmt"

	// "os"
	"strings"
)

func main() {
	saludar()
	saludar2("Mundo")
	emoji := "🍔"
	changeValue(&emoji)
	fmt.Println(emoji)
	total := suma(2, 3)
	fmt.Println(total)
	min, may := convert("Hola")
	fmt.Println(min, may)

	result, err := division(10, 4)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)

	nums := []int{1, 2, 3, 41, 25, 6, 7, 8, 90}
	res := filter(nums, func(n int) bool {
		if n >= 10 {
			return true
		}
		return false
	})

	fmt.Println(res)

	h := hello("Mundo")
	fmt.Println(h("🍔"))

	// content, err := os.ReadFile("file.txt")
	// // nil es como null en otros lenguajes
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println(string(content))

	fmt.Println(suma2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// funciones anonimas
	x := func() {
		fmt.Println("Hola")
	}
	x()

	// funciones anonimas con argumentos
	func(name string) {
		fmt.Println("Hola", name)
	}("Mundo")
}

func saludar() {
	fmt.Println("Hola")
}

func saludar2(nombre string) {
	fmt.Println("Hola", nombre)
}

func changeValue(value *string) {
	*value = "🍟"
}

// si los parametros son del mismo tipo, se puede poner el tipo al final
func suma(a, b int) int {
	return a + b
}

func convert(text string) (string, string) {
	min := strings.ToLower(text)
	may := strings.ToUpper(text)
	return min, may
}

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

func filter(nums []int, callback func(int) bool) []int {
	result := []int{}
	for _, n := range nums {
		if callback(n) {
			result = append(result, n)
		}
	}
	return result
}

func hello(name string) func(string) string {
	return func(text string) string {
		return "Hola " + name + " " + text
	}
}

func suma2(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
