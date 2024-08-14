package controller

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Pushpajit/go-bot/controller/embed"
	"github.com/Pushpajit/go-bot/utils/scrapping"
	"github.com/Pushpajit/go-bot/utils/tmdb/helper"
	"github.com/Pushpajit/go-bot/utils/tmdb/models"
	"github.com/bwmarrin/discordgo"
)

var (
	BotId   *string
	BotName *string

	deletemsg   []string
	isRecording bool = false
	mtx         sync.Mutex
	timer       *time.Timer
)

// Helper function-01
func greetings(s *discordgo.Session, m *discordgo.MessageCreate) {
	msgId, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Oh! Hi %s, nice to have you in this channel 🙂", m.Author.Username))
	if err != nil {
		log.Panic(err.Error())
	}

	// start appending the messegeID if isRecording is true
	mtx.Lock()
	if isRecording {
		mtx.Unlock()
		deletemsg = append(deletemsg, msgId.ID)
	} else {
		mtx.Unlock()
	}

}

// Helper function-02
func ghostMode(s *discordgo.Session, m *discordgo.MessageCreate, msg []string) {
	// Check if the previous scheduler is already running or not
	// then reset all the things
	if timer != nil {
		timer.Stop()
		mtx.Lock()
		isRecording = true
		deletemsg = deletemsg[:0]
		timer = nil
		mtx.Unlock()
	}

	// Make it thread safe
	mtx.Lock()
	isRecording = true
	deletemsg = append(deletemsg, m.ID)
	mtx.Unlock()

	ghostTime, _ := strconv.Atoi(msg[1]) // convert the time string to integer

	// send msg and also add this into the deletemsg slice
	msgId, _ := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("ghost mode 👻 is enabled for %v second ⏳", ghostTime))
	mtx.Lock()
	deletemsg = append(deletemsg, msgId.ID)
	mtx.Unlock()

	// scheduled function, this will run in seperate go-routine
	timer = time.AfterFunc(time.Duration(ghostTime+2)*time.Second, func() {
		mtx.Lock()
		defer mtx.Unlock()

		// send msg and also add this into the deletemsg slice
		msgId, _ := s.ChannelMessageSend(m.ChannelID, "Time's up time, auto cleaning has been started ✨🧹")
		deletemsg = append(deletemsg, msgId.ID)

		// iterate over the deletemsg and delete the messeges.
		for _, mid := range deletemsg {
			if err := s.ChannelMessageDelete(m.ChannelID, mid); err != nil {
				log.Fatalf("❌ Failed in delete msg %+#v", err.Error())
			}
		}
		// stop the ghost-mode
		isRecording = false

		// clear the deletemsg slice
		deletemsg = deletemsg[:0]

		// rest the timer
		timer = nil
	})
}

// helper function-01
func getHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	msgID, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.GetHelpEmbed())
	if err != nil {
		log.Panic(err.Error())
	}

	// start appending the messegeID if isRecording is true
	mtx.Lock()
	if isRecording {
		mtx.Unlock()
		deletemsg = append(deletemsg, msgID.ID)
	} else {
		mtx.Unlock()
	}
}

// get movies
func GetMovies(s *discordgo.Session, m *discordgo.MessageCreate, msg []string, mode int) {
	var response models.Response

	switch mode {
	case 1:
		fmt.Println("Getting GetPlayingMovie()")
		if len(msg) == 2 {
			response = helper.GetPlayingMovie(msg[1])
		} else {
			response = helper.GetPlayingMovie("")
		}

	case 2:
		fmt.Println("Getting GetPopularMovies()")
		if len(msg) == 2 {
			response = helper.GetPopularMovies(msg[1])
		} else {
			response = helper.GetPopularMovies("")
		}

	case 3:
		fmt.Println("Getting GetSearchMovie()")
		response = helper.GetSearchMovie(msg[1])

	case 4:
		fmt.Println("Getting GetSimilarMovie()")
		num, _ := strconv.Atoi(strings.Trim(msg[1], " "))
		response = helper.GetSimilarMovie(num)

	case 5:
		fmt.Println("Getting GetSuggestedMovie()")
		num, _ := strconv.Atoi(strings.Trim(msg[1], " "))
		response = helper.GetSuggestedMovie(num)

	case 6:
		fmt.Println("Getting GetUpcomingMovie()")
		if len(msg) == 2 {
			response = helper.GetUpcomingMovie(msg[1])
		} else {
			response = helper.GetUpcomingMovie("")
		}
	case 7:
		fmt.Println("Getting GetDiscoverMovie()")
		response = helper.GetDiscoverMovie(msg[1:])
	}

	result := response.Results
	// Shuffle shuffles a slice in place.
	func(slice []models.Movie) {
		rand.Seed(time.Now().UnixNano()) // Seed the random number generator
		rand.Shuffle(len(slice), func(i, j int) {
			slice[i], slice[j] = slice[j], slice[i]
		})
	}(result)

	// Send the custom created embed
	for ind, item := range result {
		if ind < 5 {
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.CreateMovieEmbed(item))
			if err != nil {
				panic(err.Error())
			}
		}
	}
}

// download images and send them to the server
func sendImage(s *discordgo.Session, m *discordgo.MessageCreate, msg []string) {
	if len(msg) < 3 {
		s.ChannelMessageSend(m.ChannelID, "syntax: !image <type of img> <count>")
		return
	}
	n, _ := strconv.Atoi(msg[2])
	urls := scrapping.GetImage(msg)

	// Shuffle the slice
	func(slice []scrapping.Image) {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(slice), func(i, j int) {
			slice[i], slice[j] = slice[j], slice[i]
		})
	}(urls)

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("✨ Generating Image Related To '%v'", msg[1]))
	for index, url := range urls {
		if index < n {
			// Create Embed
			embed := &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: url.URL,
				},
			}

			// Send Embed
			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		}
	}
}

// This function will be imported
func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == *BotId {
		return
	}

	// start appending the messegeID if isRecording is true
	mtx.Lock()
	if isRecording {
		mtx.Unlock()
		deletemsg = append(deletemsg, m.ID)
	} else {
		mtx.Unlock()
	}

	// extracting the content
	msg := strings.Split(m.Content, " ")

	// All !commands
	switch msg[0] {
	case "!hi", "!hello":
		greetings(s, m) // greet the user

	case "!ghost":
		ghostMode(s, m, msg) // enable ghost mode

	case "!help":
		getHelp(s, m) // show bot commands

	case "!image":
		sendImage(s, m, msg) // for download and send the image

	case "!movie-current":
		GetMovies(s, m, msg, 1)

	case "!movie-popular":
		GetMovies(s, m, msg, 2)

	case "!movie-search":
		GetMovies(s, m, msg, 3)

	case "!movie-similar":
		GetMovies(s, m, msg, 4)

	case "!movie-suggest":
		GetMovies(s, m, msg, 5)

	case "!movie-upcoming":
		GetMovies(s, m, msg, 6)

	case "!movie-discover":
		GetMovies(s, m, msg, 7)

	default:
		getHelp(s, m)
	}
}
