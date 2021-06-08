package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleLearn(w http.ResponseWriter, r *http.Request) {
	Logs("Learning Value")
	w.Header().Set("Content-Type", "application/json")

	var t Payload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Logs(fmt.Sprintf("Learning Payload ===> %d - %s", t.Id, t.Msg))

	BlockChain = append(BlockChain, t.Msg)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ok"))
}
