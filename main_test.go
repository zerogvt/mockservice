package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const pass = "\u2713"
const fail = "\u2717"

func init() {
	Routes()
}

func TestRootEP(t *testing.T) {
	endpoint := "/"
	t.Log("On testing endpoint" + endpoint)
	{
		testmsg := "Should be able to create a request"
		req, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			t.Fatal("\t"+testmsg, fail, err)
		}
		t.Log("\t"+testmsg, pass)

		testmsg = "Should receive 200"
		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)
		t.Log(rw.Code)
		if rw.Code != 200 {
			t.Fatal("\t"+testmsg, fail, rw.Code)
		}
		t.Log("\t"+testmsg, pass)
	}
}
