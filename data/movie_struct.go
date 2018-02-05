package data

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
