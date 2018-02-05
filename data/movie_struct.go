package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

// GetMockJSONMovie is exported because it starts with a capital letter
func GetMockJSONMovie() Movie {
	absPath, _ := filepath.Abs("../data/mock_movie.json")
	file, e := ioutil.ReadFile(absPath)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var jsontype Movie
	json.Unmarshal(file, &jsontype)
	return jsontype
}
