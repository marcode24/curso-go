package invoiceitem

// Item is a struct that represents an item in an invoice
type Item struct {
	id      uint
	product string
	value   float64
}

// New is a function that creates a new item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// ID is a method that returns the id of an item
func (i Item) Value() float64 {
	return i.value
}
