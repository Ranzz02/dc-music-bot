package handlers

import (
	"log"

	"github.com/Ranzz02/dc-music-bot/internal/commands"
	"github.com/bwmarrin/discordgo"
)

func CommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()

	switch data.Name {
	case commands.HelpCMD:
		embed, components := commands.Help(s, i) // Help command
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds:     []*discordgo.MessageEmbed{embed},
				Components: components,
				Flags:      discordgo.MessageFlagsEphemeral, // Only to sender
			},
		})
	case commands.PlayCMD:
		// Play command
		commands.Play(s, i)
	case commands.PauseCMD:
		// Pause command
		commands.Pause(s, i)
	case commands.ResumeCMD:
	// Resume command
	case commands.StopCMD:
		// Stop command
	case commands.QueueCMD:
		commands.Queue(s, i)
	default:
		log.Println("Nope")
	}

}
