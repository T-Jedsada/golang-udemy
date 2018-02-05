package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// MovieDetail is exported because it starts with a capital letter
type MovieDetail struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	Overview string `json:"overview"`
}

// Movie is exported because it starts with a capital letter
type Movie struct {
	Results []MovieDetail
}

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
	var movie Movie
	err = json.Unmarshal(body, &movie)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(movie.Results); i++ {
		fmt.Println(movie.Results[i].Title)
	}
}
