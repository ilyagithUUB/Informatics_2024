package lab4

import (
	"fmt"
	"math"
)

func calculateY(x float64) float64 {
	return math.Pow(math.Asin(x), 2) + math.Pow(math.Acos(x), 4)
}

func main() {
	fmt.Println("Задача A:")
	for x := 0.26; x <= 0.66; x += 0.08 {
		y := calculateY(x)
		fmt.Printf("x = %.2f, y = %.6f\n", x, y)
	}

	fmt.Println("\nЗадача B:")
	xValues := []float64{0.1, 0.35, 0.4, 0.55, 0.6}
	for _, x := range xValues {
		y := calculateY(x)
		fmt.Printf("x = %.2f, y = %.6f\n", x, y)
	}
}
RunLab4()