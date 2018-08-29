package main

import (
	"net/http"
	"encoding/json"
)

type Ping struct {
	Status int	`json:"status"`
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/", pingHandler)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	ping := Ping{http.StatusOK, "Hello World"}

	res, err := json.Marshal(ping)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}