package main

import "fmt"

func main() {
	/*
		if <condicion> {
			<codigo>
		} else {
			<codigo>
		}

		if <condicion> {
			<codigo>
		} else if <condicion> {
			<codigo>
		} else {
			<codigo>
		}
	*/

	emoji := "🍔"

	if emoji == "🍔" {
		fmt.Println("🍔")
	} else if emoji == "🍟" {
		fmt.Println("🍟")
	} else {
		fmt.Println("🍕")
	}

	// declarar variable en el if
	/*
		if <declaracion>; <condicion> {
			<codigo>
		}
	*/

	if emoji2 := "🍔"; emoji2 == "🍔" {
		fmt.Println("🍔")
	}
}
