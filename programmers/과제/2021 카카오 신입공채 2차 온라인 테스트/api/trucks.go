package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Trucks() trucksResponse {
	httpRequest, _ := http.NewRequest(http.MethodGet, baseUrl+"/trucks", nil)

	httpRequest.Header.Add("Content-Type", "application/json")
	httpRequest.Header.Add("Authorization", auth_key)

	client := &http.Client{}
	httpResponse, _ := client.Do(httpRequest)
	defer httpResponse.Body.Close()

	fmt.Println("response Status:", httpResponse.Status)
	fmt.Println("response Headers:", httpResponse.Header)
	body, _ := ioutil.ReadAll(httpResponse.Body)

	res := trucksResponse{}
	_ = json.Unmarshal(body, &res)
	return res
}

type trucksResponse struct {
	Trucks []struct {
		Id                int `json:"id"`
		LocationId        int `json:"location_id`
		LocatedBikesCount int `json:"located_bikes_count"`
	} `json:"trucks"`
}
