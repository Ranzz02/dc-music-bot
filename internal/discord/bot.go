package discord

import (
	"github.com/bwmarrin/discordgo"
)

func StartBot(token string) (*discordgo.Session, error) {
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return bot, nil
}
