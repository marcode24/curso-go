package greet

var spanish = "Hola"

func Greet() string {
	return spanish + " " + "Mundo"
}

// Si la primera letra de la función es mayúscula, la función es exportada
// Si la primera letra de la función es minúscula, la función es privada
// Funciona de la misma manera para las variables
