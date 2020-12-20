package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SullivanPrell/UntitledDiscordBot/discord"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Unable to load .env file")
	}

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		panic("Bot token missing - panic")
	}
	bot, err := discordgo.New("Bot " + botToken)
	if err != nil {
		panic(err)
	}

	bot.AddHandler(discord.CommandsHandlers)

	err = bot.Open()
	if err != nil {
		log.Panic("Could not connect to server")
	}

	log.Print("Bot is now running! ctrl-c to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()
}
