package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(formingMagicSquare([][]int32{{2, 9, 8}, {4, 2, 7}, {5, 6, 7}}))
}

func formingMagicSquare(s [][]int32) int32 {
	magicSquares := [][][]int32{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
	{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
	{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
	{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
	{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
	{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
	{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
	{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	}

	min := int32(100)

	for _, c := range magicSquares {
		v := calculateCost(s, c)
		if v < min {
			min = v
		}
	}

	return min
}

func calculateCost(f [][]int32, s [][]int32) int32 {
	cost := int32(0)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cost += int32(math.Abs(float64(f[i][j] - s[i][j])))
		}
	}

	return cost
}