package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/reccord/star", StartReccord).Methods("POST")
	router.HandleFunc("/api/v1/reccord/stop", StopReccord).Methods("POST")
	router.HandleFunc("/api/v1/reccord", DownloadReccord).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// StartReccord Inicia a gravação
func StartReccord(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Gravação iniciada com sucesso")
}

// StopReccord Termina a gravação
func StopReccord(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Gravação finalizada com sucesso")
}

// DownloadReccord Envia arquivo com a gravação
func DownloadReccord(w http.ResponseWriter, r *http.Request) {}
