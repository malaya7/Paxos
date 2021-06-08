package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var PORT string
var LivnessHack = false
var IngectionHack = false

func handlePromise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// jsonInfo, _ := json.Marshal(docsSlice)
	w.WriteHeader(http.StatusOK)
	//w.Write(jsonInfo)
	w.Write([]byte("Request Processed"))
}

func main() {
	port := flag.String("p", "3000", "port number")
	livens := flag.Bool("l", false, "Activate Linvness bad code")
	correctness := flag.Bool("w", false, "Acticate bad code for correctness")
	flag.Parse()

	PORT = *port
	LivnessHack = *livens
	IngectionHack = *correctness
	fmt.Println("Port====>", PORT, LivnessHack, IngectionHack)
	router := mux.NewRouter()

	router.HandleFunc("/client/propose", HandleClient).Methods("POST")
	router.HandleFunc("/learn", HandleLearn).Methods("POST")
	router.HandleFunc("/msg", HandleMsg).Methods("GET")

	if PORT == "5000" {
		if LivnessHack {
			router.HandleFunc("/prepare", HandlePrepareMalicious).Methods("POST")
			router.HandleFunc("/accept", HandleAccept).Methods("POST")
		}
		if IngectionHack {
			router.HandleFunc("/prepare", HandlePrepare).Methods("POST")
			router.HandleFunc("/accept", HandleAcceptMalicious).Methods("POST")
		}
		if !LivnessHack && !IngectionHack {
			router.HandleFunc("/prepare", HandlePrepare).Methods("POST")
			router.HandleFunc("/accept", HandleAccept).Methods("POST")
		}
	} else {
		router.HandleFunc("/prepare", HandlePrepare).Methods("POST")
		router.HandleFunc("/accept", HandleAccept).Methods("POST")
	}

	headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})

	// port := ":" + os.Getenv("PORT")
	// router.HandleFunc("/scrapers/delete/{key}", deleteHandler).Methods("GET")
	fmt.Println("server Running on port", *port)
	log.Fatal(http.ListenAndServe(":"+*port, handlers.CORS(headersOK, originsOK, methodsOK)(router)))
}
