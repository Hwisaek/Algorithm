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

/*
 * Complete the 'stringAnagram' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY dictionary
 *  2. STRING_ARRAY query
 */
func sortString(s string) string {
	a := make([]rune, len(s))
	for i, c := range s {
		a[i] = c
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return string(a)
}

func stringAnagram(dictionary []string, query []string) (a []int32) {
	a = make([]int32, len(query))
	b := make(map[string]int)

	for _, s := range dictionary {
		s2 := sortString(s)
		b[s2]++
	}

	for i, target := range query {
		target = sortString(target)
		a[i] = int32(b[target])
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

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

	result := stringAnagram(dictionary, query)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
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
