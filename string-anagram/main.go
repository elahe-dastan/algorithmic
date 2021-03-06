package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func stringAnagram(dictionary []string, query []string) []int32 {
	ans := make([]int32, 0)

	for _, q := range query {//O(q)
		counter := int32(0)
		for _, d := range dictionary {//O(d)
			if anagram(q, d) {
				counter++
			}
		}

		ans = append(ans, counter)
	}

	fmt.Println(ans)

	return ans //O(q * d * (f + s))
}

func anagram(first string, second string) bool {
	if len(first) != len(second) {
		return false
	}

	firstMap := make(map[int32]int)

	for _, f := range first {//O(f)
		v, _ := firstMap[f]
		firstMap[f] = v + 1
	}

	for _, s := range second {//O(s)
		v, ok := firstMap[s]
		if ok {
			firstMap[s] = v - 1
			if v == 0 {
				delete(firstMap, s)
			}
		}else {
			return false
		}
	}

	if len(firstMap) > 0 {
		return false
	}

	return true
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	dictionaryCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var dictionary []string

	for i := 0; i < int(dictionaryCount); i++ {
		dictionaryItem := readLine(reader)
		dictionary = append(dictionary, dictionaryItem)
	}

	queryCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var query []string

	for i := 0; i < int(queryCount); i++ {
		queryItem := readLine(reader)
		query = append(query, queryItem)
	}

	result := stringAnagram([]string{"heater", "cold", "clod", "reheat", "docl"}, []string{"codl", "heater", "abcd"})

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
