package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type dist []int32

func (f dist) Len() int {
	return len(f)
}

func (f dist) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f dist) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func getMaxArea(w int32, h int32, boundaryType []bool, boundaryDist []int32) []int64 {
	widths := dist{w}
	heights := dist{h}

	ans := make([]int64, len(boundaryType))

	for i := 0; i < len(boundaryType); i++ {
		if boundaryType[i] {
			widths = append(widths, boundaryDist[i])
			sort.Sort(widths)
		} else {
			heights = append(heights, boundaryDist[i])
			sort.Sort(heights)
		}

		w_partition := dist{}
		h_partition := dist{}

		for j := 0; j < len(widths); j++ {
			if j == 0 {
				w_partition = append(w_partition, widths[j])
			} else {
				w_partition = append(w_partition, widths[j]-widths[j-1])
			}
		}

		for j := 0; j < len(heights); j++ {
			if j == 0 {
				h_partition = append(h_partition, heights[j])
			} else {
				h_partition = append(h_partition, heights[j]-heights[j-1])
			}
		}

		sort.Sort(w_partition)
		sort.Sort(h_partition)

		ans[i] = int64(w_partition[len(w_partition)-1]) * int64(h_partition[len(h_partition)-1])
	}

	fmt.Println(ans)
	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	wTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	w := int32(wTemp)

	hTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	h := int32(hTemp)

	boundaryTypeCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var boundaryType []bool

	for i := 0; i < int(boundaryTypeCount); i++ {
		boundaryTypeItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		boundaryTypeItem := boundaryTypeItemTemp != 0
		boundaryType = append(boundaryType, boundaryTypeItem)
	}

	boundaryDistCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var boundaryDist []int32

	for i := 0; i < int(boundaryDistCount); i++ {
		boundaryDistItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		boundaryDistItem := int32(boundaryDistItemTemp)
		boundaryDist = append(boundaryDist, boundaryDistItem)
	}
	result := getMaxArea(2, 2, []bool{false, true}, []int32{1, 1})

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

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
