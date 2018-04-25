package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe("localhost:5040", nil))
	log.Fatal(http.ListenAndServe("192.168.1.5:5040", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RUL.Path = %q\n", r.URL.Path)
}
