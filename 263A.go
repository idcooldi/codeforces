package main

import (
	"fmt"
	"math"
)

func main() {
	matrix1 := make([]int, 5)
	matrix2 := make([][]int, 5)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&matrix1[j])
		}
		matrix2[i] = append(matrix2[i], matrix1...)
	}
	var row, column int
	for k, v := range matrix2 {
		for x, y := range v {
			if y == 1 {
				row = x
				column = k
			}
		}
	}
	fmt.Println(math.Abs(2-float64(row)) + math.Abs(2-float64(column)))
}
