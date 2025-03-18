package bot

import (
	"log"

	"github.com/Ranzz02/dc-music-bot/config"
	"github.com/Ranzz02/dc-music-bot/internal/commands"
	"github.com/Ranzz02/dc-music-bot/internal/discord"
	"github.com/Ranzz02/dc-music-bot/internal/handlers"
	"github.com/bwmarrin/discordgo"
)

func Start() {
	conf := config.NewEnvConfig()
	bot, err := discord.StartBot(conf.Token)
	if err != nil {
		panic(err)
	}

	// Register interaction handlers
	bot.AddHandler(handlers.CommandHandler)

	bot.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = bot.Open()
	if err != nil {
		log.Fatalf("Error opening Discord connection: %e", err)
	}

	// Set bots status
	err = bot.UpdateGameStatus(-1, conf.Prefix+"help")
	if err != nil {
		log.Fatalf("Error updating status %v", err)
	}

	// Register commands (change GUILD_ID or use "" for global)
	commands.RegisterCommands(bot, "")

	log.Println("🤖 Bot is running... Press CTRL+C to exit.")

	select {}
}
