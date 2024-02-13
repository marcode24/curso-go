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

	// switch
	/*
		switch <expresion> {
		case <valor>:
			<codigo>
		case <valor>:
			<codigo>
		default:
			<codigo>
		}
	*/

	fmt.Println("Switch")

	// no es necesario poner break al final de cada case
	// si se quiere que no siga evaluando los siguientes case, se pone fallthrough
	// fallthrough

	switch emoji {
	case "🍔":
		fmt.Println("🍔")
	case "🍟":
		fmt.Println("🍟")
	default:
		fmt.Println("🍕")
	}

	// switch con fallthrough
	switch emoji {
	case "🍔":
		fmt.Println("🍔")
		fallthrough
	case "🍟":
		fmt.Println("🍟")
	default:
		fmt.Println("🍕")
	}
}
