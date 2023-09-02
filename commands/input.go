package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func InputRun(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	name := options[0].StringValue()
	major := options[1].BoolValue()
	var majorString string

	if major {
		majorString = "Yes"
	} else {
		majorString = "No"
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Name: " + name + "\nMajor: " + majorString,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func InputRegister() *discordgo.ApplicationCommand {
	dm := false
	return &discordgo.ApplicationCommand{
		Name:         "input",
		Description:  "Input!",
		DMPermission: &dm,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Your name",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "major",
				Description: "Are you major",
				Required:    true,
			},
		},
	}
}
