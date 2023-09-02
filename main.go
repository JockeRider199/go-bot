package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/JockeRider199/go-bot/commands"
	"github.com/JockeRider199/go-bot/events"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	registerCommands = []*discordgo.ApplicationCommand{commands.PingRegister(), commands.InputRegister()}
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	token := os.Getenv("CLIENT_TOKEN")
	guildId := os.Getenv("GUILD_ID")

	client, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating client.")
	}

	client.AddHandler(func(*discordgo.Session, *discordgo.Ready) {
		log.Println("Bot is up!")
	})
	client.AddHandler(events.MessageCreate)
	client.AddHandler(events.InteractionCreate)

	client.Identify.Intents |= discordgo.IntentsGuildMessages
	// session.Identify.Intents |= discordgo.Intent...

	err = client.Open()
	if err != nil {
		log.Fatal("Error creating client connection.")
	}

	_, err = client.ApplicationCommandBulkOverwrite(client.State.User.ID, guildId, registerCommands)
	if err != nil {
		log.Fatalf("Error creating command %s", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	client.Close()
}
