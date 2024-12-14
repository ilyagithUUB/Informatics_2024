package lab4

import (
	"fmt"
	"math"
)

func calculateY(x float64) float64 {
	return math.Pow(math.Asin(x), 2) + math.Pow(math.Acos(x), 4)
}

func TaskA(xn, xk, deltax float64) [][]float64 {
	var taskA [][]float64
	for x := xn; x <= xk; x += deltax {
		y := calculateY(x)
		taskA = append(taskA, []float64{x, y})
	}
	return taskA
}

func TaskB(values []float64)[][]float64 {
	var taskB [][]float64
	for _, x := range values {
		y := calculateY(x)
		taskB = append(taskB, []float64{x, y})
	}
	return taskB
}

func RunLab4() {
	resultsA := TaskA(0.2, 0.95, 0.15)
	for _, result := range resultsA {
		fmt.Printf("x: %.2f, y: %.2f\n", result[0], result[1])
	}

	arr := []float64{0.1, 0.35, 0.4, 0.55, 0.6}
	resultsB := TaskB(arr)
	for _, result := range resultsB {
		fmt.Printf("x: %.2f, y: %.2f\n", result[0], result[1])
	}
}