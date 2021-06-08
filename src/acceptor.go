package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var BlockChain = []string{"Genisis"}

func HandleAccept(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle Accept", IngectionHack)
	if IngectionHack {
		HandleAcceptMalicious(w, r)
		return
	}
	Logs("Accept Request")
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
	ress := fmt.Sprintf("Ack Accepted \"%s\"", t.Msg)
	Logs(ress)

	BlockChain = append(BlockChain, t.Msg)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Accepted - Normal"))
}

func HandleMsg(w http.ResponseWriter, r *http.Request) {
	Logs("GeT Msg")
	w.Header().Set("Content-Type", "application/json")
	lastVal := BlockChain[len(BlockChain)-1]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(lastVal))
}
