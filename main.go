package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var host *string = flag.String("host", "localhost", "server host, ip")
var port *int = flag.Int("port", 8000, "the default port")

func main() {
	flag.Parse()

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
