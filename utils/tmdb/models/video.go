package models

type Video struct {
	Type     string `json:"type"`
	Site     string `json:"site"`
	Name     string `json:"name"`
	Official bool   `json:"official"`
	Key      string `json:"key"`
}

type MovieResponse struct {
	Results []Video `json:"results"`
}
