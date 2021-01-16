package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(100)
	if v > 90 {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	io.WriteString(w, "hello")
}

func client() {
	for {
		time.Sleep(1 * time.Second)
		r, err := http.Get("http://localhost:8000/")
		if err == nil {
			fmt.Println(r.StatusCode)
		} else {
			fmt.Println(err)
		}
	}
}

func main() {
	fmt.Println("hello")
	http.HandleFunc("/", rootHandler)
	go client()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
