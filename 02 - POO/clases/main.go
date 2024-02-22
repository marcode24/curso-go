package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

func (c Course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

func main() {
	Go := Course{
		Name:    "Go desde cero",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	// Otra forma de declarar una estructura
	// se debe declarar todos los campos
	// teniendo el mismo orden que la estructura
	Go2 := Course{
		"Go desde cero",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	fmt.Print(Go.Name + "\n")
	fmt.Print(Go2.Name + "\n")

	// PrintClasses(Go)
	Go.PrintClasses()
	(&Go).ChangePrice(15.5)
	fmt.Print(Go.Price)
}
