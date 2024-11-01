package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed :(", http.StatusMethodNotAllowed)

		return
	}

	dat := Response{
		Text:      "The backend is now on. 後端已連接。",
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// or else the site won't be able to fetch it

	json.NewEncoder(w).Encode(dat)
}

func main() {
	http.HandleFunc("/", reqHandler)

	fmt.Println("The server is now running at <http://localhost:8080> :)")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
