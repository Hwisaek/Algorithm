package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Start(problem int) startResponse {
	startReq := startRequest{problem}
	bodyBytes, _ := json.Marshal(startReq)

	req, _ := http.NewRequest(http.MethodPost, baseUrl+"/start", bytes.NewBuffer(bodyBytes))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", x_auth_token)

	client := &http.Client{}
	response, _ := client.Do(req)
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)

	startRes := startResponse{}
	_ = json.Unmarshal(body, &startRes)

	auth_key = startRes.AuthKey
	return startRes
}

type startRequest struct {
	Problem int `json:"problem"`
}

type startResponse struct {
	AuthKey string `json:"auth_key"`
	Problem int    `json:"problem"`
	Time    int    `json:"time"`
}
