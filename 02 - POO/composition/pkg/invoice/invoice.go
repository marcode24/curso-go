package invoice

import (
	"github.com/marcode24/composition/pkg/customer"
	"github.com/marcode24/composition/pkg/invoiceitem"
)

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   []invoiceitem.Item
}

// New is a function that creates a new invoice
func New(country, city string, client customer.Customer, items []invoiceitem.Item) Invoice {
	return Invoice{country, city, 0, client, items}
}

// SetTotal is a method that sets the total of an invoice
func (i *Invoice) SetTotal() {
	i.total = 0
	for _, item := range i.items {
		i.total += item.Value()
	}
}
