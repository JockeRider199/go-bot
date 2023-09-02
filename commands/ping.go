package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func PingRun(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func PingRegister() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Ping!",
	}
}
