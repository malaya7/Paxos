package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var numReq = map[string]int{"3000": 0, "5000": 0, "8000": 0}
var c = 0

func HandlePrepareMidleware(w http.ResponseWriter, r *http.Request) {
	Logs("Propose- Middlewaere")
	w.Header().Set("Content-Type", "application/json")
	s := strings.Split(r.Host, ":")
	fmt.Println("sssssssssssssssss", s)
	if s[1] == "5000" {
		c += 1
	}
	if c >= 2 {
		Logs("Node acting maliciously")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Promise"))
	}

	var t Payload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Logs(fmt.Sprintf("Proposeal Payload ===> %d - %s", t.Id, t.Msg))
	if t.Id <= PromiseVal {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed, Id lower than expected"))
		Logs("Failed, Id lower than expected")
		return
	}
	PromiseVal = t.Id
	Logs(fmt.Sprintf("Promised value \"%s\"", t.Msg))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Promise"))
}

func HandleAcceptMidleware(w http.ResponseWriter, r *http.Request) {
	Logs("Accept Request - Middleware")
	w.Header().Set("Content-Type", "application/json")

	var t Payload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Logs(fmt.Sprintf("Proposeal Payload ===> %d - %s", t.Id, t.Msg))

	if t.Id != PromiseVal {
		w.WriteHeader(http.StatusBadRequest)
		er := fmt.Sprintf("Other Value %d Already Promised", t.Id)
		w.Write([]byte(er))
		Logs(er)
		return
	}

	t.Id++
	t.Msg = fmt.Sprintf("GET Hacked %d", t.Id)

	json_data, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf("%s/learn", n3Url)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))
	if resp.StatusCode != 200 {
		Logs("Protocol Faile,d Node 3 did not learn")
	}
	Logs("Protocol Failed, Ingected bad value to node 2 and node 3")

	BlockChain = append(BlockChain, t.Msg)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Accepted"))
}
