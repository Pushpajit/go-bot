package bot

import (
	"fmt"
	"log"

	"github.com/Pushpajit/go-bot/config"
	"github.com/Pushpajit/go-bot/controller"
	"github.com/bwmarrin/discordgo"
)

var (
	BotId   string
	BotName string
)

func StartBOT() {
	// start a new session for bot
	goBot, err := discordgo.New(fmt.Sprintf("Bot %s", config.Auth.Token))

	// check for error during session creation
	if err != nil {
		log.Fatalf("❌ Bot session failed: %v", err.Error())
		return
	}

	// Get the BotID and assign
	uid, err := goBot.User("@me")
	if err != nil {
		log.Fatalf("❌ Bot user-fetch failed: %v", err.Error())
	}

	// Set the BotId and BotName to the controller
	controller.BotId = &uid.ID
	controller.BotName = &uid.GlobalName

	// Add a bot handler function, basically a controller for bot commands
	goBot.AddHandler(controller.Handler)

	// now start the bot, make it online
	if err := goBot.Open(); err != nil {
		log.Fatalf("❌ Bot start failed: %v", err.Error())
	}

}
