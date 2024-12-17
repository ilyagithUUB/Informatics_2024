package lab7

import "fmt"

type Electronics struct {
	Name        string
	Price       float64
	Brand       string
	Description string
}

func (e *Electronics) GetName() string {
	return e.Name
}

func (e *Electronics) GetPrice() float64 {
	return e.Price
}

func (e *Electronics) SetPrice(price float64) {
	e.Price = price
}

func (e *Electronics) ApplyDiscount(discount float64) {
	e.Price -= e.Price * discount
}

func (e *Electronics) GetInfo() string {
	return fmt.Sprintf("Электроника: %s, Бренд: %s, Цена: %.2f", e.Name, e.Brand, e.Price)
}