package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	var server = fmt.Sprintf("%s:%d", *host, *port)
	fmt.Println("will Serve on:", server)

	setHandler()
	log.Fatal(http.ListenAndServe(server, nil))

	fmt.Printf("Serve on %s success! But cannot reach here\n", server)
}

func setHandler() {
	http.HandleFunc("/url", handler)
	http.HandleFunc("/", welcome)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, wrold!")
	fmt.Fprintln(w, "Hello, golang!")
	fmt.Fprintln(w, "Hello, vim note!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RUL.Path = %q\n", r.URL.Path)
}
