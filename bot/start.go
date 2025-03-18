package bot

import (
	"github.com/Ranzz02/dc-music-bot/config"
	"github.com/Ranzz02/dc-music-bot/internal/discord"
)

func Start() {
	conf := config.NewEnvConfig()
	bot, err := discord.StartBot(conf.Token)
	if err != nil {
		panic(err)
	}
	bot.Open()
}
