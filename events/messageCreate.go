package events

import "github.com/bwmarrin/discordgo"

func MessageCreate(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Content == "!ping" {
		s.ChannelMessageSend(msg.ChannelID, "Pong!")
	}
}
