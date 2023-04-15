package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Simulate() simulateResponse {
	startReq := simulateRequest{Commands: []command{{
		TruckId: 0,
		Command: []int{2, 5, 4, 1, 6},
	}}}
	bodyBytes, _ := json.Marshal(startReq)

	httpRequest, _ := http.NewRequest(http.MethodPut, baseUrl+"/simulate", bytes.NewBuffer(bodyBytes))

	httpRequest.Header.Add("Content-Type", "application/json")
	httpRequest.Header.Add("Authorization", auth_key)

	client := &http.Client{}
	httpResponse, _ := client.Do(httpRequest)
	defer httpResponse.Body.Close()

	fmt.Println("response Status:", httpResponse.Status)
	fmt.Println("response Headers:", httpResponse.Header)
	body, _ := ioutil.ReadAll(httpResponse.Body)

	res := simulateResponse{}
	_ = json.Unmarshal(body, &res)
	return res
}

type command struct {
	TruckId int   `json:"truck_id"`
	Command []int `json:"command"`
}
type simulateRequest struct {
	Commands []command `json:"commands"`
}

type simulateResponse struct {
	Status              string  `json:"status"`
	Time                int     `json:"time"`
	FailedRequestsCount int     `json:"failed_requests_count"`
	Distance            float64 `json:"distance"`
}
