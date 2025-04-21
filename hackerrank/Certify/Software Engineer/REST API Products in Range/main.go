package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getProductsInRange' function below.
 *
 * URL for cut and paste
 * https://jsonmock.hackerrank.com/api/inventory?category=<category>&page=<pageNumber>
 *
 * The function is expected to return an Integer value.
 * The function accepts String category, Integer minPrice and Integer maxPrice as arguments.
 *
 */

func getProductsInRange(category string, minPrice int32, maxPrice int32) int32 {
	return 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	category := readLine(reader)

	minPriceTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	minPrice := int32(parentIdTemp)

	maxPriceTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	maxPrice := int32(maxPriceTemp)

	result := getProductsInRange(category, minPrice, maxPrice)

	fmt.Fprintf(writer, "%d\n", result)

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
