package lab4

import (
	"fmt"
	"math"
)

func Calculate(x float64) float64 {
	y := math.Pow(math.Pow(math.Asin(x), 2)+math.Pow(math.Acos(x), 4), 3)
	return y
}

func main() {
	xValues := []float64{0.1, 0.35, 0.4, 0.55, 0.6}
	for _, x := range xValues {
		y := math.Pow(math.Pow(math.Asin(x), 2)+math.Pow(math.Acos(x), 4), 3)
		fmt.Printf("x = %.2f, y = %.6f\n", x, y)
	}
}
