package events

import (
	"github.com/JockeRider199/go-bot/commands"
	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "ping":
		commands.PingRun(s, i)
	case "input":
		commands.InputRun(s, i)
	}
}
