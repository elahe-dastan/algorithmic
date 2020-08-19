package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the bomberMan function below.
type Grid struct {
	IsBomb bool
}

type Bomb struct {
	Row    int
	Column int
}

func bomberMan(n int32, grid []string) []string {
	if n % 2 == 0 {
		res := make([]string, len(grid))

		for i := 0; i < len(grid); i++ {
			row := grid[i]
			r := ""

			for j := 0; j < len(row); j++ {
				r += "O"
			}

			res[i] = r
		}

		return res
	}

	g := make([][]Grid, len(grid))

	for i := range g {
		row := grid[i]
		r := make([]Grid, len(row))
		for j := 0; j < len(row); j++ {
			if row[j] == 'O' {
				r[j] = Grid{IsBomb: true}
			} else {
				r[j] = Grid{IsBomb: false}
			}
		}

		g[i] = r
	}

	n-- // first
	if n == 0 {
		return gridToString(g)
	}

	for {
		n-- // second
		bombs := make([]Bomb, 0)
		for i, r := range g {
			for j := 0; j < len(r); j++ {
				if g[i][j].IsBomb {
					bombs = append(bombs, Bomb{
						Row:    i,
						Column: j,
					})
				}
			}
		}

		for i, s := range g {
			for j := 0; j < len(s); j++ {
				g[i][j].IsBomb = true
			}
		}

		n-- // third

		for _, b := range bombs {
			g[b.Row][b.Column].IsBomb = false // itself

			if b.Row != 0 {
				g[b.Row-1][b.Column].IsBomb = false
			} // up

			if b.Row != len(g)-1 {
				g[b.Row+1][b.Column].IsBomb = false
			} //down

			if b.Column != 0 {
				g[b.Row][b.Column-1].IsBomb = false
			} //left

			if b.Column != len(g[0])-1 {
				g[b.Row][b.Column+1].IsBomb = false
			} //right
		}

		if n == 0 {
			return gridToString(g)
		}
	}
}

func gridToString(g [][]Grid) []string {
	res := make([]string, len(g))

	for i := 0; i < len(g); i++ {
		row := g[i]
		r := ""

		for j := 0; j < len(row); j++ {
			if row[j].IsBomb {
				r += "O"
			}else {
				r += "."
			}
		}

		res[i] = r
	}

	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	rcn := strings.Split(readLine(reader), " ")

	rTemp, err := strconv.ParseInt(rcn[0], 10, 64)
	checkError(err)
	r := int32(rTemp)

	_, err = strconv.ParseInt(rcn[1], 10, 64)
	checkError(err)
	nTemp, err := strconv.ParseInt(rcn[2], 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(r); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := bomberMan(n, grid)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result) - 1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
