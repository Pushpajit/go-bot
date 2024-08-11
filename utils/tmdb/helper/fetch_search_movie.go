package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Pushpajit/go-bot/utils/tmdb/models"
	"github.com/joho/godotenv"
)

func GetSearchMovie(query string) models.Response {
	var response models.Response

	if query == "" {
		return response
	}

	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("❌ ENV load failed: %v", err.Error())
	}

	// endpoint to fetch the popular
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?query=%v&include_adult=false&language=en-US&page=1", strings.Join(strings.Split(query, "-"), "%20"))

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("MOVIETOKEN")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		fmt.Println("Error:", err)
	}

	return response
}
