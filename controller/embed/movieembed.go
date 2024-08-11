package embed

import (
	"fmt"
	"strings"

	"github.com/Pushpajit/go-bot/utils/tmdb/models"
	"github.com/bwmarrin/discordgo"
)

// CreateMovieEmbed generates a Discord embed message for a given movie
func CreateMovieEmbed(movie models.Movie) *discordgo.MessageEmbed {
	// Base URL for images
	baseImageURL := "https://image.tmdb.org/t/p/original"

	// Make fields for genres
	var formattedGenres []string
	for _, genreID := range movie.Genres {
		formattedGenres = append(formattedGenres, fmt.Sprintf("`%s`", models.MovieGenre[genreID]))
	}

	generateGenres := strings.Join(formattedGenres, ", ")

	fmt.Println("generateGenres: ", generateGenres)

	return &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("🎬 **%s** `(ID: %d)`", movie.Title, movie.Id),
		Description: fmt.Sprintf("📝 **Overview**:\n%s", movie.Overview),
		Color:       0x00b4d8, // Golang color
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "⭐ Rating",
				Value:  fmt.Sprintf("%.1f / 10", movie.Rating),
				Inline: true,
			},
			{
				Name:   "📅 Release Date",
				Value:  movie.Date,
				Inline: true,
			},
			{
				Name:   "🎭 Genres",
				Value:  generateGenres,
				Inline: false,
			},
			{
				Name:   "🔗 More Info",
				Value:  fmt.Sprintf("[TMDb Link](https://www.themoviedb.org/movie/%d)", movie.Id),
				Inline: false,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: baseImageURL + movie.Backdrop,
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: baseImageURL + movie.Poster,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Powered by Golang 🔵",
			IconURL: "https://www.themoviedb.org/assets/2/v4/logos/208x226-dark-bg-3c327a981ea5acb36252467403d230e06a8f878b453d39d2c1f75b88b22bcd08.png",
		},
	}
}
