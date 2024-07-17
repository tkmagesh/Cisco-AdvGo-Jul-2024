package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

// struct composition
type PerishableProduct struct {
	Product
	Expiry string
}

// method overriding
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

// fmt.Stringer interface implementation
func (pp PerishableProduct) String() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

// factory function to hide the complexity of object construction
func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}
func main() {
	pen := Product{Id: 100, Name: "Pen", Cost: 10}
	fmt.Println(pen.Format())
	pen.ApplyDiscount(10)
	fmt.Println(pen.Format())

	// Perishable Product
	milk := PerishableProduct{
		Product: Product{
			Id:   200,
			Name: "Milk",
			Cost: 45,
		},
		Expiry: "1 Day",
	}
	fmt.Println("milk.Product.Id =", milk.Product.Id)
	fmt.Println("milk.Id =", milk.Id) // attribute inheritance

	// method inheritance
	// fmt.Println(milk.Format())
	fmt.Println(milk)
	milk.ApplyDiscount(10)
	// fmt.Println(milk.Format())
	fmt.Println(milk)

	grapes := NewPerishableProduct(201, "Grapes", 100, "2 Days")
	// fmt.Println(grapes.Format())
	fmt.Println(grapes)
}
