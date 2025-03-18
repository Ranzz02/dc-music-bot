package commands

import "github.com/bwmarrin/discordgo"

func Play(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSend(m.ChannelID, "Playing music!")
}
