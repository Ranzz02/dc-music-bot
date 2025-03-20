package commands

import (
	"log"
	"os/exec"
	"strconv"
	"sync"

	"github.com/bwmarrin/discordgo"
)

// Song represents a song in the queue
type Song struct {
	URL         string
	User        string
	Interaction *discordgo.InteractionCreate
	Session     *discordgo.Session
}

// SongQueue manages the song queue
type SongQueue struct {
	Songs []Song
	mu    sync.Mutex // ✅ Fix: Initialize the mutex
}

// Guild song queues
var guildQueues = make(map[string]*SongQueue)

// Add a song to the queue
func AddToQueue(guildID, userID, songUrl string, interaction *discordgo.InteractionCreate, session *discordgo.Session) {
	queue, exists := guildQueues[guildID]
	if !exists {
		queue = &SongQueue{Songs: []Song{}} // ✅ Initialize the queue properly
		guildQueues[guildID] = queue
	}

	queue.mu.Lock()
	defer queue.mu.Unlock()

	queue.Songs = append(queue.Songs, Song{URL: songUrl, User: userID, Interaction: interaction, Session: session})
	log.Printf("Added song to queue: %s", songUrl)
}

// Play the next song in the queue
func PlayNextSong(vc *discordgo.VoiceConnection, guildID string) {
	queue, exists := guildQueues[guildID]
	if !exists || len(queue.Songs) == 0 {
		log.Println("Queue is empty, leaving voice channel.")
		vc.Disconnect()
		return
	}

	// Get the next song in the queue
	song := queue.Songs[0]
	log.Printf("Playing song from queue: %s", song.URL)

	// Play the song asynchronously
	go playYouTubeAudio(vc, song, guildID)
}

// Clean up the queue after playback
func CleanUpQueue(guildID string) {
	delete(guildQueues, guildID)
}

// Play YouTube audio
func playYouTubeAudio(vc *discordgo.VoiceConnection, song Song, guildID string) {
	if vc == nil {
		log.Println("Voice connection is nil, cannot play audio.")
		return
	}

	cmd := exec.Command("yt-dlp", "-f", "bestaudio", "--no-playlist", "-q", "-o", "-", song.URL)
	cmd.Stderr = log.Writer()

	ffmpeg := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "s16le", "-ac", "2", "-ar", "48000", "pipe:1")
	ffmpeg.Stdin, _ = cmd.StdoutPipe()
	ffmpeg.Stderr = log.Writer()

	// Run the Node.js Opus encoder
	opusCmd := exec.Command("node", "encoder/index.js")
	opusCmd.Stdin, _ = ffmpeg.StdoutPipe()
	opusCmd.Stderr = log.Writer()
	opusPipe, err := opusCmd.StdoutPipe()
	if err != nil {
		log.Printf("Error creating Opus encoder stdout pipe: %v", err)
		return
	}

	// Start yt-dlp
	err = cmd.Start()
	if err != nil {
		log.Printf("Error starting yt-dlp: %v", err)
		return
	}

	// Start ffmpeg convertion
	err = ffmpeg.Start()
	if err != nil {
		log.Printf("Error starting ffmpeg: %v", err)
		return
	}

	// Start opus encoding
	err = opusCmd.Start()
	if err != nil {
		log.Printf("Error starting Node.js Opus encoder: %v", err)
		return
	}

	vc.Speaking(true)
	defer vc.Speaking(false)

	buf := make([]byte, 960*2*2)
	for {
		n, err := opusPipe.Read(buf)
		if err != nil {
			log.Printf("Finished playing song: %v", err)
			break
		}

		log.Printf("Read bytes: %v", n)
		vc.OpusSend <- buf[:n]
	}

	cmd.Wait()
	ffmpeg.Wait()
	opusCmd.Wait()

	// Remove the song from the queue
	queue, exists := guildQueues[guildID]
	if exists {
		queue.mu.Lock()
		queue.Songs = queue.Songs[1:] // Remove the first song
		queue.mu.Unlock()
	}

	// Play the next song
	PlayNextSong(vc, guildID)
}

// Queue command setup
var QueueCommand *discordgo.ApplicationCommand = &discordgo.ApplicationCommand{
	Name:        "queue", // ✅ Fix: Properly define the command name
	Description: "List current music queue",
}

// Queue command handler
func Queue(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guildID := i.GuildID
	queue, exists := guildQueues[guildID]
	if !exists || len(queue.Songs) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Queue is currently empty",
			},
		})
		return
	}

	var songInfos []*discordgo.MessageEmbedField
	for idx, song := range queue.Songs {
		songInfos = append(songInfos, &discordgo.MessageEmbedField{
			Value: strconv.Itoa(idx+1) + " " + song.URL,
		})
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Current queue",
		Description: "List of all songs in queue",
		Color:       0x3498db,
		Fields:      songInfos,
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}
