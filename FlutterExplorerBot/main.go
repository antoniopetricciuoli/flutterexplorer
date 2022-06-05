package main

import (
	"flutterexplorerbot/handlers"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	client := gobotapi.NewClient(os.Getenv("BOT_TOKEN"))

	client.OnCommand("start", []string{"/", ";", "."}, handlers.HandleStart)
	client.OnInlineQuery(handlers.HandleInline)

	client.Run()
}
