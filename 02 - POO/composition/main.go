package main

import (
	"fmt"

	"github.com/marcode24/composition/pkg/customer"
	"github.com/marcode24/composition/pkg/invoice"
	"github.com/marcode24/composition/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"USA",
		"New York",
		customer.New("John Doe", "123 C. Street", "555-5555"),
		[]invoice.Item{
			invoiceitem.New(1, "Laptop", 1000),
			invoiceitem.New(2, "Mouse", 10),
		},
	)

	i.SetTotal()
	fmt.Printf("%+v\n", i)
}
