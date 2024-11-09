package lab4

import (
	"fmt"
	"math"
)

func calculateY(x float64) float64 {
	return math.Pow(math.Asin(x), 2) + math.Pow(math.Acos(x), 4)
}

func TaskA() {
	fmt.Println("Задача A:")
	for x := 0.26; x <= 0.66; x += 0.08 {
		y := calculateY(x)
		fmt.Printf("x = %.2f, y = %.6f\n", x, y)
	}
}
func TaskB() {
	fmt.Println("\nЗадача B:")
	xValues := []float64{0.1, 0.35, 0.4, 0.55, 0.6}
	for _, x := range xValues {
		y := calculateY(x)
		fmt.Printf("x = %.2f, y = %.6f\n", x, y)
	}
}

func RunLab4(){
	resultsA := TaskA(0.26, 0.66, 0.08)
  for _, result := range resultsA {
    fmt.Printf("x = %.2f, y = %.6f\n", x, y)
  }

  arr := []float64{0.1, 0.35, 0.4, 0.55, 0.6}
  resultsB := TaskB( 0.26, 0.66, arr)
  for _, result := range resultsB {
    fmt.Printf("x = %.2f, y = %.6f\n", x, y)
}
}