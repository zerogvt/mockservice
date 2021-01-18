package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Settings is our
type Settings struct {
	ErrorRate     int `json:"err"`
	ReqsPerMinute int `json:"rpm"`
}

var defSet = Settings{
	ErrorRate:     10,
	ReqsPerMinute: 600}

var settings *Settings

func init() {
	rand.Seed(time.Now().UnixNano())
	settings = &defSet
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	v := rand.Intn(100)
	if v < settings.ErrorRate {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	io.WriteString(w, "hello")
}

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	v, ok := q["err"]
	if ok {
		settings.ErrorRate, _ = strconv.Atoi(v[0])
	}
	v, ok = q["rpm"]
	if ok {
		settings.ReqsPerMinute, _ = strconv.Atoi(v[0])
	}
	io.WriteString(w, fmt.Sprintf("error rate: %d, requests per minute: %d",
		settings.ErrorRate, settings.ReqsPerMinute))
}

func client() {
	for {
		sleepT := time.Millisecond * time.Duration(60000/settings.ReqsPerMinute)
		time.Sleep(sleepT)
		r, err := http.Get("http://localhost:8000/")
		if err == nil {
			fmt.Println(r.StatusCode)
		} else {
			fmt.Println(err)
		}
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/set/", settingsHandler)
	go client()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
