package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const (
	PlayCMD   string = "play"
	PauseCMD  string = "pause"
	ResumeCMD string = "resume"
	StopCMD   string = "stop"
)

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        PlayCMD,
		Description: "Play a song from YouTube",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "query",
				Description: "YouTube URL or search query",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	},
	{
		Name:        "hello",
		Description: "Hello, world!",
	},
	{
		Name:        StopCMD,
		Description: "Stop the music and leave the voice channel",
	},
}

func RemoveAllCommands(s *discordgo.Session, guildID string) {
	commands, err := s.ApplicationCommands(s.State.User.ID, guildID)
	if err != nil {
		log.Printf("Error fetching commands for guild %s: %v", guildID, err)
	}

	// Loop through all the commands and delete them individually
	for _, cmd := range commands {
		err := s.ApplicationCommandDelete(s.State.User.ID, guildID, cmd.ID)
		if err != nil {
			log.Printf("Error removing command %s from guild %s: %v", cmd.Name, guildID, err)
		} else {
			log.Printf("Successfully removed command %s from guild %s.", cmd.Name, guildID)
		}
	}
}

func RegisterCommands(s *discordgo.Session, guildID string) {
	for _, cmd := range Commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, cmd)
		if err != nil {
			log.Fatalf("Cannot create command %q: %v", cmd.Name, err)
		}
	}
	log.Printf("✅ all %d/%d commands registered!", len(Commands), len(Commands))
}
