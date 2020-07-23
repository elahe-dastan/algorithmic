package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var vowels = []uint8{'a', 'e', 'i', 'o', 'u'}

func findSubstring(s string, k int32) string {
	var i int32
	max := int32(0)
	best := "Not found!"
	maxNumOfVowels := int32(0)
	counter := int32(0)

	for i := 0; i < len(s); i++ {//O(n)
		for _, v := range vowels {//5
			if s[i] == v {
				counter++
				break
			}
		}

		if counter == k {
			break
		}
	}

	if counter == 0 {
		return best
	}

	maxNumOfVowels = counter

	var v int32
	for i = 0; i <= int32(len(s))-k; i++ { //O(n - k)
		if i == 0 {
			v = countVowels(s[i : i+k])
		}else {
			v = nextSub(s[i-1 : i+k], v)
		}
		if v == maxNumOfVowels {
			return s[i : i+k]
		}

		if v > max {
			max = v
			best = s[i : i+k]
		}
	}

	fmt.Println(best)

	return best //O(5 * n) + O(n - k) * 5
}

func countVowels(s string) int32 {
	counter := int32(0)

	for i := 0; i < len(s); i++ {//O(k)
		for _, v := range vowels {//5
			if s[i] == v {
				counter++
				break
			}
		}
	}

	return counter
}

func nextSub(s string, count int32) int32 {
	for _, v := range vowels {//5
		if s[0] == v {
			count--
			break
		}
	}

	for _, v := range vowels {//5
		if s[len(s) - 1] == v {
			count++
			break
		}
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := findSubstring("azerdii", 5)

	fmt.Fprintf(writer, "%s\n", result)

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
