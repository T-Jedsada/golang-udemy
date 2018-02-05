package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-udemy/data"
	"github.com/gorilla/handlers"
	_ "github.com/heroku/x/hmetrics/onload"
)

var (
	httpPort = os.Getenv("PORT")
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
	if httpPort == "" {
		httpPort = "8080"
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/movie", handlerMovie)
	http.ListenAndServe(":"+httpPort, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
