package commands

import (
	"github.com/bwmarrin/discordgo"
)

func Play(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	s.ChannelMessageSend(m.ChannelID, "Playing music!")

	_, err := s.ChannelVoiceJoin(m.GuildID, m.ChannelID, false, true)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Failed to join voice channel.")
	}

}
