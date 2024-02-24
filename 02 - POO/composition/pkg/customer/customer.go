package customer

// Customer is a struct that represents a customer
type Customer struct {
	name    string
	address string
	phone   string
}

// New is a function that creates a new customer
func New(name, address, phone string) Customer {
	return Customer{name, address, phone}
}
