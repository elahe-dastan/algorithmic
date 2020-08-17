package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(encryption("iffactsdontfittotheorychangethefacts"))
}

func encryption(s string) string {
	floor := math.Floor(math.Sqrt(float64(len(s))))
	ceil := math.Ceil(math.Sqrt(float64(len(s))))

	if floor * floor >= float64(len(s)) {
		mtx := make([][]uint8, int(floor))
		counter := 0

		for i := 0; i < int(floor); i++ {
			row := make([]uint8, int(floor))
			for j := 0; j < int(floor); j++ {
				if counter >= len(s) {
					row[j] = 0
					continue
				}
				row[j] = s[counter]
				counter++
			}
			mtx[i] = row
		}

		return encrypt(floor, floor, mtx)
	}else if floor * ceil >= float64(len(s)){
		mtx := make([][]uint8, int(floor))
		counter := 0

		for i := 0; i < int(floor); i++ {
			row := make([]uint8, int(ceil))
			for j := 0; j < int(ceil); j++ {
				if counter >= len(s) {
					row[j] = 0
					continue
				}
				row[j] = s[counter]
				counter++
			}
			mtx[i] = row
		}

		return encrypt(floor, ceil, mtx)
	}else {
		mtx := make([][]uint8, int(ceil))
		counter := 0

		for i := 0; i < int(ceil); i++ {
			row := make([]uint8, int(ceil))
			for j := 0; j < int(ceil); j++ {
				if counter >= len(s) {
					row[j] = 0
					continue
				}
				row[j] = s[counter]
				counter++
			}
			mtx[i] = row
		}

		return encrypt(ceil, ceil, mtx)
	}
}

func encrypt(floor float64, ceil float64, mtx [][]uint8) string {
	var encrypted []uint8

	for k := 0; k < int(ceil); k++ {
		for l := 0; l < int(floor); l++ {
			if mtx[l][k] != uint8(0) {
				encrypted = append(encrypted, mtx[l][k])
			}
		}

		encrypted = append(encrypted, uint8(' '))
	}

	return string(encrypted)
}
