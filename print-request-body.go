package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func headerBody(rw http.ResponseWriter, r *http.Request) {
	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}

func test() {
	http.HandleFunc("/", headerBody)
	http.ListenAndServe(":8080", nil)
}
