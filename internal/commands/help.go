package commands

import (
	"github.com/Ranzz02/dc-music-bot/config"
	"github.com/bwmarrin/discordgo"
)

var conf config.EnvConfig = *config.NewEnvConfig()

var (
	HelpCommand *discordgo.ApplicationCommand = &discordgo.ApplicationCommand{
		Name:        HelpCMD,
		Description: "List all commands",
	}
)

const (
	GeneralHelpCMD string = "general"
	MusicHelpCMD   string = "music"
	CloseHelpCMD   string = "close"
)

func Help(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.MessageEmbed, []discordgo.MessageComponent) {
	embed := &discordgo.MessageEmbed{
		Title:       "📖 Help Menu",
		Description: "Here are the available commands:",
		Color:       0x3498db, // Blue color
		Fields: []*discordgo.MessageEmbedField{
			{Name: conf.Prefix + HelpCMD, Value: "Help command, lists and helps with command use", Inline: false},
			{Name: conf.Prefix + PlayCMD, Value: "Play a song from YouTube", Inline: false},
			{Name: conf.Prefix + StopCMD, Value: "Stop the current song and remove from queue", Inline: false},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Use the buttons below to explore more!",
		},
	}

	components := []discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					CustomID: GeneralHelpCMD,
					Label:    "General Commands",
					Style:    discordgo.PrimaryButton,
				},
				&discordgo.Button{
					CustomID: MusicHelpCMD,
					Label:    "Music Commands",
					Style:    discordgo.SuccessButton,
				},
				&discordgo.Button{
					CustomID: CloseHelpCMD,
					Label:    "Close",
					Style:    discordgo.DangerButton,
				},
			},
		},
	}

	return embed, components
}
