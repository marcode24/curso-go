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

	// for
	/*
		for <inicializacion>; <condicion>; <incremento> {
			<codigo>
		}
	*/

	fmt.Println("For")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for continuo
	/*
		for <condicion> {
			<codigo>
		}
	*/

	fmt.Println("For continuo")

	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	// for forever
	/*
		for {
			<codigo>
		}
	*/

	// for range
	/*
		for <indice>, <valor> := range <arreglo> {
			<codigo>
		}
	*/

	fmt.Println("For range")

	nums := []uint8{2, 4, 6, 8, 10}

	for i, v := range nums {
		fmt.Println("Indice:", i, "Valor:", v)
	}

	for i := range nums {
		nums[i] *= 2
	}
	fmt.Println(nums)

	sports := map[string]string{
		"basketball": "🏀",
		"football":   "⚽",
		"baseball":   "⚾",
	}

	for k, v := range sports {
		fmt.Println(k, v)
	}

	hello := "Hello"
	for i, v := range hello {
		fmt.Println(i, string(v), v)
	}
}
