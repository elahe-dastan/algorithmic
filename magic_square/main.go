package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(formingMagicSquare([][]int32{{2, 9, 8}, {4, 2, 7}, {5, 6, 7}}))
}

func formingMagicSquare(s [][]int32) int32 {
	magicSquares := permutations([]int32{9, 8, 7, 6})

	min := int32(100)

	for _, c := range magicSquares {
		v := calculateCost(s, c)
		if v < min {
			min = v
		}
	}

	return min
}

func calculateCost(f [][]int32, s []int32) int32 {
	firstMin := math.Abs(float64(f[0][0]-s[0])) + math.Abs(float64(f[2][2]-(10-s[0])))
	if math.Abs(float64(f[2][2]-s[0])) + math.Abs(float64(f[0][0]-(10-s[0]))) < firstMin {
		firstMin = math.Abs(float64(f[2][2]-s[0])) + math.Abs(float64(f[0][0]-(10-s[0])))
	}

	secondMin := math.Abs(float64(f[0][1]-s[1])) + math.Abs(float64(f[2][1]-(10-s[1])))
	if math.Abs(float64(f[2][1]-s[1])) + math.Abs(float64(f[0][1]-(10-s[1]))) < secondMin {
		secondMin = math.Abs(float64(f[2][1]-s[1])) + math.Abs(float64(f[0][1]-(10-s[1])))
	}

	thirdMin := math.Abs(float64(f[0][2]-s[2])) + math.Abs(float64(f[2][0]-(10-s[2])))
	if math.Abs(float64(f[2][0]-s[2])) + math.Abs(float64(f[0][2]-(10-s[2]))) < thirdMin {
		thirdMin = math.Abs(float64(f[2][0]-s[2])) + math.Abs(float64(f[0][2]-(10-s[2])))
	}

	fourthMin := math.Abs(float64(f[1][0]-s[3])) + math.Abs(float64(f[1][2]-(10-s[3])))
	if math.Abs(float64(f[1][2]-s[3])) + math.Abs(float64(f[1][0]-(10-s[3]))) < fourthMin {
		fourthMin = math.Abs(float64(f[1][2]-s[3])) + math.Abs(float64(f[1][0]-(10-s[3])))
	}

	fifthMin := math.Abs(float64(f[1][1] - 5))

	return int32(firstMin + secondMin + thirdMin + fourthMin + fifthMin)
}
func permutations(arr []int32)[][]int32{
	var helper func([]int32, int32)
	var res [][]int32

	helper = func(arr []int32, n int32){
		if n == 1{
			tmp := make([]int32, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := int32(0); i < n; i++{
				helper(arr, n - 1)
				if n % 2 == 1{
					tmp := arr[i]
					arr[i] = arr[n - 1]
					arr[n - 1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n - 1]
					arr[n - 1] = tmp
				}
			}
		}
	}
	helper(arr, int32(len(arr)))
	return res
}