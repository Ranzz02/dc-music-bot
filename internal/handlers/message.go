package handlers

import (
	"strings"

	"github.com/Ranzz02/dc-music-bot/config"
	"github.com/Ranzz02/dc-music-bot/internal/commands"
	"github.com/bwmarrin/discordgo"
)

func CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	conf := config.NewEnvConfig()

	// Check if message starts with the prefix
	if !strings.HasPrefix(m.Content, conf.Prefix) {
		return
	}

	// Remove the prefix to process the command
	command := strings.TrimPrefix(m.Content, conf.Prefix)
	args := strings.Fields(command)

	switch command {
	case commands.PlayCMD:
		commands.Play(s, m, args)
	case "hello":
		s.ChannelMessageSend(m.ChannelID, "World!")
	}
}
