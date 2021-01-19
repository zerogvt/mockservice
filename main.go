package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// Settings are our operational parameters
type Settings struct {
	ErrorRate     int `json:"err"`
	ReqsPerMinute int `json:"rpm"`
}

var defSet = Settings{
	ErrorRate:     10,
	ReqsPerMinute: 30}

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

// Routes sets our public endpoints
func Routes() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("mockservice_go"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
	)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", rootHandler))
	http.HandleFunc("/set/", settingsHandler)
}

func main() {
	Routes()
	go client()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
