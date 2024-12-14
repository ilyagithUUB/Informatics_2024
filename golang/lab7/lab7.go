package lab7

import "fmt"

type Product interface {
	ApplyDiscount(discount float64)
	SetPrice(newPrice float64)
	GetPrice() float64
	getInfo() string
}

func GetTotalPrice(products []Product) float64 {
	var TotalPrice float64 = 0
	for _, Product := range products {
		TotalPrice += Product.GetPrice()
	}
	return TotalPrice
}

func RunLab7() {
	Product1 := &Electronics{"Washer", 5000.0, "Bosch", "Bosch"}
	Product2 := &Book{"Harry Potter and the Half-Blood Prince", 500.0, "Black-blue-Yellow", "Black-blue-Yellow"}
	fmt.Println(Product1.GetDescription())
	fmt.Println(Product2.GetDescription())
}