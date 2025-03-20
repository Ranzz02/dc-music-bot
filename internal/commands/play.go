package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	PlayCommand *discordgo.ApplicationCommand = &discordgo.ApplicationCommand{
		Name:        PlayCMD,
		Description: "Play a song from YouTube",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "query",
				Description: "YouTube URL or search query",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        "immediate",
				Description: "Skip queue straight to this song",
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Required:    false,
			},
		},
	}
)

func Play(s *discordgo.Session, i *discordgo.InteractionCreate) {
	query := i.ApplicationCommandData().Options[0].StringValue()

	guildID := i.GuildID
	voiceState, err := s.State.VoiceState(guildID, i.Member.User.ID)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Needs to be in a voice channel to use.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	// Add to queue
	AddToQueue(guildID, i.Member.User.ID, query, i, s)

	// Try to join voice channel
	vc, err := s.ChannelVoiceJoin(i.GuildID, voiceState.ChannelID, false, true)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Failed to join voice channel.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	// If this is the first song, start playing it
	if len(guildQueues[guildID].Songs) >= 1 {
		// Play the first song in the queue
		PlayNextSong(vc, guildID)
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "🎶 Playing: " + query,
		},
	})
}
