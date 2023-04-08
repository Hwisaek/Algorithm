package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Score() scoreResponse {
	httpRequest, _ := http.NewRequest(http.MethodGet, baseUrl+"/score", nil)

	httpRequest.Header.Add("Content-Type", "application/json")
	httpRequest.Header.Add("Authorization", auth_key)

	client := &http.Client{}
	httpResponse, _ := client.Do(httpRequest)
	defer httpResponse.Body.Close()

	fmt.Println("response Status:", httpResponse.Status)
	fmt.Println("response Headers:", httpResponse.Header)
	body, _ := ioutil.ReadAll(httpResponse.Body)

	res := scoreResponse{}
	_ = json.Unmarshal(body, &res)
	return res
}

type scoreRequest struct {
	Commands []struct {
		TruckId int      `json:"truck_id"`
		Command []string `json:"command"`
	} `json:"commands"`
}

type scoreResponse struct {
	Score float64 `json:"score"`
}
