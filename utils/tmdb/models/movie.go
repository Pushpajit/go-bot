package models

type Movie struct {
	Id       int     `json:"id"`
	Title    string  `json:"title"`
	Overview string  `json:"overview"`
	Rating   float64 `json:"vote_average"`
	Date     string  `json:"release_date"`
	Backdrop string  `json:"backdrop_path"`
	Poster   string  `json:"poster_path"`
	Genres   []int   `json:"genre_ids"`
}

type Response struct {
	Results []Movie `json:"results"`
}

var MovieGenre = map[int]string{
	28:    "Action",
	12:    "Adventure",
	16:    "Animation",
	35:    "Comedy",
	80:    "Crime",
	99:    "Documentary",
	18:    "Drama",
	10751: "Family",
	14:    "Fantasy",
	36:    "History",
	27:    "Horror",
	10402: "Music",
	9648:  "Mystery",
	10749: "Romance",
	878:   "Science Fiction",
	10770: "TV Movie",
	53:    "Thriller",
	10752: "War",
	37:    "Western",
}

var Catagory = map[string]int{
	"Action":          28,
	"Adventure":       12,
	"Animation":       16,
	"Comedy":          35,
	"Crime":           80,
	"Documentary":     99,
	"Drama":           18,
	"Family":          10751,
	"Fantasy":         14,
	"History":         36,
	"Horror":          27,
	"Music":           10402,
	"Mystery":         9648,
	"Romance":         10749,
	"Science Fiction": 878,
	"TV Movie":        10770,
	"Thriller":        53,
	"War":             10752,
	"Western":         37,
}
