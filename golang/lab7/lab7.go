package lab7

import "fmt"

type Product interface {
	ApplyDiscount(discount float64)
	SetPrice(newPrice float64)
	GetPrice() float64
	GetInfo() string
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
	
	ProductsWithoutDiscount := []Product{Product1, Product2}
	
	fmt.Printf("Информация о продуктах:\n")
	for _, p := range ProductsWithoutDiscount {
		fmt.Println(p.GetInfo())
	}
	
	fmt.Printf("\nЦена без скидки: %.2f рублей\n", GetTotalPrice(ProductsWithoutDiscount))
	
	Product1.ApplyDiscount(0.10)
	Product2.ApplyDiscount(0.20)
	
	fmt.Printf("Цена со скидкой: %.2f рублей\n", GetTotalPrice(ProductsWithoutDiscount))
	
	fmt.Printf("\nИнформация о продуктах после скидки:\n")
	for _, p := range ProductsWithoutDiscount {
		fmt.Println(p.GetInfo())
	}
}