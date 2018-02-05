package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang-udemy/data"
)

func main() {
	resp, err := http.Get("https://workshopup.herokuapp.com/movie")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var movie data.Movie
	err = json.Unmarshal(body, &movie)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(movie.Results); i++ {
		fmt.Println(movie.Results[i].Title)
	}
}
