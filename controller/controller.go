package controller

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

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
	embed := &discordgo.MessageEmbed{
		Title:       "Bot Commands!",
		Description: "Here are the available commands:",
		Color:       0x6AD7E4,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "!hi or !hello",
				Value:  "Replies with a greeting message. ✨🎉🎊",
				Inline: false,
			},
			{
				Name:   "!ghost <seconds>",
				Value:  "Enables ghost mode 👻 for the specified duration in seconds. Messages sent during this time will be deleted automatically after the time expires.",
				Inline: false,
			},
			{
				Name:   "!help ",
				Value:  "Show all the available bot commands. 🆘",
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Working on more commands 🎮, please stay tuned! 🎆🎇",
		},
	}

	msgID, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
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
	}
}
