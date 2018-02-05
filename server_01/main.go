package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-udemy/data"
)

var (
	httpPort = ":8080"
	listenIP = "localhost"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Fucking Shit Man!!!")
}

func handlerMovie(w http.ResponseWriter, r *http.Request) {
	movie, err := json.Marshal(data.GetMockJSONMovie())
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(movie)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/movie", handlerMovie)
	http.ListenAndServe(httpPort, nil)
}
