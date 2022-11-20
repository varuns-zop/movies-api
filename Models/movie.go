package Models

type MovieDetails struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   *Movie `json:"data"`
}

type Movie struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Rating   string `json:"rating"`
	Plot     string `json:"plot"`
	Released bool   `json:"released"`
}
