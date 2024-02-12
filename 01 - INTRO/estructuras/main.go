package main

import "fmt"

func main() {
	type course struct {
		Name string
		Slug string
	}

	db := course{
		Name: "Curso Profesional de Go",
		Slug: "curso-profesional-go",
	}

	fmt.Printf("%+v\n", db)
	fmt.Println(db)

	// se puede omitir el nombre del campo
	// solo se puede hacer si se inicializa todos los campos
	git := course{
		"Git desde cero",
		"git-desde-cero",
	}

	fmt.Printf("%+v\n", git)
	fmt.Println(git)

	// para acceder a un campo
	fmt.Println(db.Name)
	fmt.Println(git.Slug)

	p := &db
	(*p).Name = "Curso Profesional de Go 2"

	fmt.Printf("%+v\n", p)
}
