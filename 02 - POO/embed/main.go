package main

import "fmt"

// Person is a struct that represents a person
type Person struct {
	Name string
	Age  uint
}

// NewPerson is a function that creates a new person
func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

// Greet is a method that prints a greeting
func (p Person) Greet() {
	fmt.Println("Hello")
}

// Human is a struct that represents a human
type Human struct {
	Age      uint
	Children uint
}

// NewHuman is a function that creates a new human
func NewHuman(age, children uint) Human {
	return Human{age, children}
}

// Employee is a struct that represents an employee
type Employee struct {
	Person
	Human
	Salary float64
}

// Greet is a method that prints a greeting
func (e Employee) Greet() {
	fmt.Println("Hola desde empleado")
}

// NewEmployee is a function that creates a new employee
func NewEmployee(name string, age, children uint, salary float64) Employee {
	return Employee{NewPerson(name, age), NewHuman(age, children), salary}
}

// Payroll is a method that prints the payroll
func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.90)
}

func main() {
	e := NewEmployee("Alejandro", 30, 2, 1000000)
	fmt.Println(e.Name, e.Human.Age)
	e.Person.Greet()
	e.Greet()
	e.Payroll()
}
