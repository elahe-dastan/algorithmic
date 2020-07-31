package main

import "fmt"

func main() {
	fmt.Println(isValidSudoku([][]byte{
	{'5','3','8','.','7','.','.','.','.'},
	{'6','.','.','1','9','5','.','.','.'},
	{'.','9','8','.','.','.','.','6','.'},
	{'8','.','.','.','6','.','.','.','3'},
	{'4','.','.','8','.','3','.','.','1'},
	{'7','.','.','.','2','.','.','.','6'},
	{'.','6','.','.','.','.','2','8','.'},
	{'.','.','.','4','1','9','.','.','5'},
	{'.','.','.','.','8','.','.','7','9'}}))
}

func isValidSudoku(board [][]byte) bool {
	for _, row := range board {
		if checkRepetition(row) {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		column := make([]byte, 9)

		for j, row := range board {
			column[j] = row[i]
		}

		if checkRepetition(column) {
			return false
		}
	}

	for i := 0; i < 9; i = i+3 {
		box := make([]byte, 9)
		f := board[i]
		s := board[i+1]
		t := board[i+2]
		for j := 0; j < 9; j= j+3 {
			box[0] = f[j]
			box[1] = f[j+1]
			box[2] = f[j+2]

			box[3] = s[j]
			box[4] = s[j+1]
			box[5] = s[j+2]

			box[6] = t[j]
			box[7] = t[j+1]
			box[8] = t[j+2]

			if checkRepetition(box) {
				return false
			}
		}
	}

	return true
}

func checkRepetition(arr []byte) bool {
	m := make(map[byte]bool)

	for _, b := range arr {
		if b != byte('.') {
			_, ok := m[b]
			if ok {
				return true
			}
			m[b]= true
		}
	}

	return false
}
