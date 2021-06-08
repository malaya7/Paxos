package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var RoundID int = 0
var PromiseVal = 0

const node2Port = 5000
const node3Port = 8000

var node2Url string = fmt.Sprintf("http://localhost:%d", node2Port)
var node3Url string = fmt.Sprintf("http://localhost:%d", node3Port)

var Nodes = []string{node2Url, node3Url}

type Payload struct {
	Id  int
	Msg string
}

func HandleClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	RoundID++
	Logs(fmt.Sprintf("New Client Request, RoundId: %d", RoundID))

	var t Payload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// log.Println("Propose Payload ===>", t)

	p := Payload{Id: RoundID, Msg: t.Msg}

	for _, v := range Nodes {
		json_data, err := json.Marshal(p)
		if err != nil {
			log.Fatal(err)
		}
		url := fmt.Sprintf("%s/prepare", v)
		Logs("ProposeTo==>" + url)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != 200 {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("Failed, Unable to Promise"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Promised"))
}

func HandlePrepare(w http.ResponseWriter, r *http.Request) {
	Logs("Recivied New Propose")
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
