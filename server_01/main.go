package main

import (
	"fmt"
	"net/http"
)

var (
	httpPort = ":8080"
	listenIP = "localhost"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handlerMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Movie.json")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/movie", handlerMovie)
	http.ListenAndServe(httpPort, nil)
}
