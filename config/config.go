package config

import (
	"log"
	"os"

)

type BotConfig struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botprefix"`
}

var (
	Auth *BotConfig
)

func ReadConfig() error {
	// load .env file
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("❌ ENV load failed: %v", err.Error())
	// }

	// set the Auth variable
	Auth = &BotConfig{
		Token:     os.Getenv("TOKEN"),
		BotPrefix: os.Getenv("BOTPREFIX"),
	}

	return nil
}
