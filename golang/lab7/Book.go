package lab7

import "fmt"

type Book struct {
	Name        string
	Price       float64
	Author      string
	Description string
}

func (b *Book) GetName() string {
	return b.Name
}

func (b *Book) GetPrice() float64 {
	return b.Price
}

func (b *Book) SetPrice(price float64) {
	b.Price = price
}

func (b *Book) ApplyDiscount(discount float64) {
	b.Price -= b.Price * discount
}

func (b *Book) GetInfo() string {
	return fmt.Sprintf("Книга: %s, Автор: %s, Цена: %.2f", b.Name, b.Author, b.Price)
}