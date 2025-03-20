package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	PlayCMD   string = "play"
	SkipCMD   string = "skip"
	PauseCMD  string = "pause"
	ResumeCMD string = "resume"
	StopCMD   string = "stop"
	QueueCMD  string = "queue"
	HelpCMD   string = "help"
)

var Commands = []*discordgo.ApplicationCommand{
	PlayCommand,  // Play command: adds to queue and plays song
	QueueCommand, // Queue command: lists queue
	HelpCommand,  // Help command: lists all commands
}

func RegisterCommands(s *discordgo.Session, guildID string) {
	created, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, guildID, Commands)
	if err != nil {
		log.Fatalf("Error registering commands: %v", err)
		return
	}
	log.Printf("✅  %d/%d commands created or overwritten!", len(created), len(Commands))
}
