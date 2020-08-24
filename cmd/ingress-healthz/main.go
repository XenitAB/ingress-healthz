package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ResponseData struct {
	Time           time.Time
	Env            string
	RequestHeaders http.Header
}

func handler(w http.ResponseWriter, r *http.Request) {
	// log.Print("ingress-healthz received a request.")

	environmentShort := os.Getenv("ENV")
	if environmentShort == "" {
		environmentShort = "lab"
	}

	responseData := ResponseData{
		Time:           time.Now(),
		Env:            environmentShort,
		RequestHeaders: r.Header,
	}

	jsonResponse, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(string(jsonResponse))

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResponse); err != nil {
		log.Println(err)
	}
}

func main() {
	log.Print("ingress-healthz started.")

	http.HandleFunc("/healthz", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
