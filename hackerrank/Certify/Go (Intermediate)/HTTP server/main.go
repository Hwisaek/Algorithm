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
	"time"
)

/*
 * Complete the 'postHandler', 'deleteHandler' and 'getHandler' functions below.
 *
 * All functions are expected to be void.
 * All functions accept http.ResponseWriter w and *http.Request req as parameters.
 */

func postHandler(w http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	var l Lake
	json.Unmarshal(body, &l)
	store[l.Id] = l
}

func deleteHandler(w http.ResponseWriter, req *http.Request) {
	id := strings.Split(strings.Split(strings.Split(req.RequestURI, "?")[1], "&")[0], "=")[1]
	delete(store, id)
}

func getHandler(w http.ResponseWriter, req *http.Request) {
	id := strings.Split(strings.Split(strings.Split(req.RequestURI, "?")[1], "&")[0], "=")[1]
	l, ok := store[id]
	if !ok {
		w.WriteHeader(404)
		return
	}

	b, _ := json.Marshal(l)

	w.WriteHeader(200)
	w.Write(b)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	actionsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/delete", deleteHandler)
	go http.ListenAndServe(portSuffix, nil)
	time.Sleep(100 * time.Millisecond)

	var actions []string

	for i := 0; i < int(actionsCount); i++ {
		actionsItem := readLine(reader)
		actions = append(actions, actionsItem)
	}

	for _, actionStr := range actions {
		var action Action
		err := json.Unmarshal([]byte(actionStr), &action)
		checkError(err)
		switch action.Type {
		case "post":
			_, err := http.Post(address+"/post", "application/json", strings.NewReader(action.Payload))
			checkError(err)
		case "delete":
			client := &http.Client{}
			req, err := http.NewRequest("DELETE", address+"/delete?id="+action.Payload, nil)
			checkError(err)
			resp, err := client.Do(req)
			checkError(err)
			if resp.StatusCode != 200 {
				fmt.Fprintf(writer, "%s\n", resp.Status)
				continue
			}
		case "get":
			resp, err := http.Get(address + "/get?id=" + action.Payload)
			checkError(err)
			if resp.StatusCode != 200 {
				fmt.Fprintf(writer, "%s\n", resp.Status)
				continue
			}
			var lake Lake
			err = json.NewDecoder(resp.Body).Decode(&lake)
			checkError(err)
			fmt.Fprintf(writer, "%s\n", lake.Name)
			fmt.Fprintf(writer, "%d\n", lake.Area)
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

const portSuffix = ":3333"

var address = "http://127.0.0.1" + portSuffix

type Lake struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Area int32  `json:"area"`
}

type Action struct {
	Type    string
	Payload string
}

var store = map[string]Lake{}

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
