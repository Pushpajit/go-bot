package embed

import "github.com/bwmarrin/discordgo"

func GetHelpEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "🛠️ **Bot Command Help**",
		Description: "Here are all the commands you can use with this bot:",
		Color:       0x00ffcc, // A cool teal color

		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "👋 `!hi` / `!hello`",
				Value:  "Greet the bot and it will greet you back!",
				Inline: false,
			},
			{
				Name:   "👻 `!ghost`",
				Value:  "Enable ghost mode. **Syntax:** `!ghost`",
				Inline: false,
			},
			{
				Name:   "❓ `!help`",
				Value:  "Display this help message. **Syntax:** `!help`",
				Inline: false,
			},
			{
				Name:   "🖼️ `!image`",
				Value:  "Download and send an image. **Syntax:** `!image [keyword] [number] [format]`",
				Inline: false,
			},
			{
				Name:   "🎬 `!movie-current`",
				Value:  "Fetch current movies. **Syntax:** `!movie-current`",
				Inline: false,
			},
			{
				Name:   "🔥 `!movie-popular`",
				Value:  "Fetch popular movies. **Syntax:** `!movie-popular`",
				Inline: false,
			},
			{
				Name:   "🔍 `!movie-search`",
				Value:  "Search for a movie. **Syntax:** `!movie-search [query]`",
				Inline: false,
			},
			{
				Name:   "🔗 `!movie-similar`",
				Value:  "Find movies similar to a given one. **Syntax:** `!movie-similar [movie ID]`",
				Inline: false,
			},
			{
				Name:   "🎲 `!movie-suggest`",
				Value:  "Get movie suggestions based on genre. **Syntax:** `!movie-suggest [genre]`",
				Inline: false,
			},
			{
				Name:   "🎥 `!movie-upcoming`",
				Value:  "Fetch upcoming movies. **Syntax:** `!movie-upcoming`",
				Inline: false,
			},
			{
				Name:   "🌐 `!movie-discover`",
				Value:  "Discover movies based on various criteria. **Syntax:** `!movie-discover [options]`",
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Use the commands wisely! 😉",
		},
	}
}
