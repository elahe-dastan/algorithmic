package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the climbingLeaderboard function below.
func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	scoreSet := make([]int32, 0)

	for i, v := range scores{
		if i == 0{
			scoreSet = append(scoreSet, v)
		}else {
			if v != scores[i - 1] {
				scoreSet = append(scoreSet, v)
			}
		}
	}

	res := make([]int32, len(alice))
	for j, v := range alice {
		if j == 0 {
			res[j] = rank(scoreSet, v)
		}else {
			var max int32
			if int32(len(scoreSet)) < res[j - 1] {
				max = int32(len(scoreSet))
			}else {
				max = res[j - 1]
			}
			res[j] = rank(scoreSet[:max], v)
		}
	}

	return res
}

func rank(scoreSet []int32, score int32) int32 {
	low := 0
	high := len(scoreSet)

	for low < high {
		middle := (low + high) / 2

		if scoreSet[middle] > score {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}

	if low == len(scoreSet) {
		return int32(low + 1)
	} else if scoreSet[low] <= score {
		return int32(low + 1)
	} else {
		return int32(low + 2)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	scoresTemp := strings.Split(readLine(reader), " ")

	var scores []int32

	for i := 0; i < int(scoresCount); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	aliceTemp := strings.Split(readLine(reader), " ")

	var alice []int32

	for i := 0; i < int(aliceCount); i++ {
		aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
		checkError(err)
		aliceItem := int32(aliceItemTemp)
		alice = append(alice, aliceItem)
	}

	result := climbingLeaderboard(scores, alice)

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
