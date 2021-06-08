package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const node3 = 8000

var n3Url string = fmt.Sprintf("http://localhost:%d", node3)

func HandlePrepareMalicious(w http.ResponseWriter, r *http.Request) {
	Logs("Recivied New Propose - Act Malicious")
	w.Header().Set("Content-Type", "application/json")

	var t Payload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if LivnessHack || t.Id <= PromiseVal {
		t.Id++
		t.Msg = "Hacking Liveness"
		PromiseVal = t.Id

		json_data, err := json.Marshal(t)
		if err != nil {
			log.Fatal(err)
		}

		url := fmt.Sprintf("%s/prepare", n3Url)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(json_data))

		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode != 200 {
			Logs("Protocol Failed by overwriting Most recent RoundId")
		}
		Logs("Protocol Failed -  Act Malicious proposed own RoundId")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed, Id lower than expected"))
		Logs("Failed, Id lower than expected")
		return
	}

	PromiseVal = t.Id
	t.Msg = "GET Hacked"
	Logs(fmt.Sprintf("Injected bad value \"%s\"", t.Msg))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Promise"))
}

func HandleAcceptMalicious(w http.ResponseWriter, r *http.Request) {
	Logs("Accept Request - Act Malicious")
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

	// make the other node learn bad value while telling the proposers accepted Original value
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
	w.Write([]byte("Accepted - Act malici"))
}
