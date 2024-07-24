package main

import (
	"fmt"
	"log"

	"github.com/Pushpajit/go-bot/bot"
	"github.com/Pushpajit/go-bot/config"
)

func main() {
	fmt.Println("Hi!, I'm go-bot 🤖")

	if err := config.ReadConfig(); err != nil {
		log.Fatalf("❌ Bot start failed: %v", err.Error())
	}

	fmt.Println("🟢 Bot is online")
	bot.StartBOT()

	<-make(chan struct{})

}
