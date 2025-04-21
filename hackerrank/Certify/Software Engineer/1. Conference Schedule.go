package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
type Response struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []data `json:"data"`
}

type data struct {
	Barcode   string  `json:"barcode"`
	Item      string  `json:"item"`
	Category  string  `json:"category"`
	Price     int32   `json:"price"`
	Discount  float64 `json:"discount"`
	Available int     `json:"available"`
}

func getProductsInRange(category string, minPrice int32, maxPrice int32) (a int32) {
	_, _, _, totalPages, _, err := request(category, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range totalPages {
		_, _, _, _, dataList, err := request(category, i)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, d := range dataList {
			if d.Price >= minPrice && d.Price <= maxPrice {
				a++
			}
		}
	}

	return
}

func request(category string, reqPage int) (page, perPage, total, totalPages int, data []data, err error) {
	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/inventory?category=%s&page=%d", category, reqPage)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	return response.Page, response.PerPage, response.Total, response.TotalPages, response.Data, nil
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	category := readLine(reader)

	minPriceTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	minPrice := int32(minPriceTemp)

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
