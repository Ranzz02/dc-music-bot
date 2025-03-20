package handlers

import (
	"github.com/Ranzz02/dc-music-bot/internal/commands"
	"github.com/bwmarrin/discordgo"
)

func HelpButtonHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	backButton := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					CustomID: "back",
					Label:    "Back",
					Style:    discordgo.DangerButton,
				},
			},
		},
	}

	switch i.MessageComponentData().CustomID {
	case commands.GeneralHelpCMD: // General commands
		embed := &discordgo.MessageEmbed{
			Title:       "📖 General Commands",
			Description: "Here are some general commands:",
			Color:       0xf1c40f, // Yellow color
			Fields: []*discordgo.MessageEmbedField{
				{Name: "/help", Value: "Shows this help menu", Inline: false},
				{Name: "/ping", Value: "Check bot latency", Inline: false},
			},
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage, // Updates the message
			Data: &discordgo.InteractionResponseData{
				Embeds:     []*discordgo.MessageEmbed{embed},
				Components: backButton,
			},
		})
	case commands.MusicHelpCMD: // Music commands help
		embed := &discordgo.MessageEmbed{
			Title:       "🎵 Music Commands",
			Description: "Here are the available music commands:",
			Color:       0x2ecc71, // Green color
			Fields: []*discordgo.MessageEmbedField{
				{Name: "/play", Value: "Play a song from YouTube", Inline: false},
				{Name: "/stop", Value: "Stop the current song", Inline: false},
				{Name: "/skip", Value: "Skip the current song", Inline: false},
				{Name: "/queue", Value: "Show the current queue", Inline: false},
			},
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Embeds:     []*discordgo.MessageEmbed{embed},
				Components: backButton,
			},
		})
	case commands.CloseHelpCMD: // Close help menu
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Content:    "Help menu closed.",
				Embeds:     nil,
				Components: nil, // Remove buttons
			},
		})
	case "back":
		embed, components := commands.Help(s, i)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Embeds:     []*discordgo.MessageEmbed{embed},
				Components: components,
			},
		})
	}
}
