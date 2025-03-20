import {
	AudioPlayerStatus,
	createAudioPlayer,
	createAudioResource,
	joinVoiceChannel,
	NoSubscriberBehavior,
	StreamType,
	VoiceConnection,
	VoiceConnectionStatus,
} from "@discordjs/voice";
import { CommandInteraction, EmbedBuilder, GuildMember } from "discord.js";
import { stream, validate, video_basic_info } from "play-dl";

const MAX_DURATION = 600;

interface Song {
	url: string;
	title: string;
	duration: number;
}

const queue: Song[] = [];

let connection: VoiceConnection | null = null;
let isPlaying = false;

// Main function
export async function Play(interaction: CommandInteraction) {
	const member = interaction.member as GuildMember;
	if (!member.voice.channel) {
		const embed = errorEmbed(
			"",
			"You must be in a voice channel to play music!"
		);
		return interaction.reply({
			embeds: [embed],
			flags: "Ephemeral",
		});
	}

	// Respond with first message
	interaction.deferReply();

	// Get the query (YouTube URL)
	const query = interaction.options.get("query", true)?.value as string;

	// Validate YouTube URL
	if ((await validate(query)) !== "yt_video") {
		const embed = errorEmbed(
			query,
			"Invalid YouTube URL. Please provide a valid YouTube link."
		);
		return interaction.reply({
			embeds: [embed],
			flags: "Ephemeral",
		});
	}

	// Get video info
	try {
		const info = await video_basic_info(query);
		const duration = info.video_details.durationInSec;

		// Enforce max duration limit
		if (duration > MAX_DURATION) {
			const embed = errorEmbed(
				query,
				`The video is too long! Maximum allowed length is ${
					MAX_DURATION / 60
				} minutes.`
			);
			return interaction.reply({
				embeds: [embed],
				flags: "Ephemeral",
			});
		}

		const song: Song = {
			url: query,
			title: info.video_details.title || "",
			duration: duration,
		};

		queue.push(song);

		if (!connection) {
			connection = joinVoiceChannel({
				channelId: member.voice.channelId!,
				guildId: interaction.guildId!,
				adapterCreator: member.guild.voiceAdapterCreator,
			});

			connection.on(VoiceConnectionStatus.Disconnected, () => {
				connection?.destroy();
				connection = null;
			});
		}

		if (!isPlaying) {
			playSong(interaction);
			await interaction.reply({
				content: `Now playing: **${song.title}**`,
			});
		} else {
			await interaction.reply({
				content: `✅ Added to queue: **${song.title}**`,
			});
		}
	} catch (error) {
		console.log(error);
		const embed = errorEmbed(
			query,
			"There was an error fetching the video details. Please try again."
		);
		return interaction.reply({
			embeds: [embed],
			flags: "Ephemeral",
		});
	}
}

// Go through queue
async function playSong(interaction: CommandInteraction) {
	if (queue.length === 0) {
		isPlaying = false;
		connection?.destroy();
		return;
	}

	const song = queue[0];

	let player = createAudioPlayer({
		behaviors: {
			noSubscriber: NoSubscriberBehavior.Play,
		},
	});
	try {
		let audio_stream = await stream(song.url, {});
		console.log("🔊 Stream created:", stream.name);

		let resource = createAudioResource(audio_stream.stream, {
			inputType: audio_stream.type,
		});

		console.log(`▶️ Now playing: ${song.title}`);
		player.play(resource);

		console.log("🔊 Subscribing player to connection...");
		connection?.subscribe(player)!;
		console.log("✅ Player subscribed!");

		player.on(AudioPlayerStatus.Idle, () => {
			// When the song ends, play the next song in the queue
			queue.shift(); // Remove the current song from the queue
			playSong(interaction);
		});

		player.on(AudioPlayerStatus.Playing, () => {
			console.log("🎵 Music is now playing...");
		});

		player.on("error", (error) => {
			console.log(error);
		});
	} catch (error) {
		console.log("Error playing song: ", error);
		queue.shift();
		playSong(interaction);
	}
}

function errorEmbed(query: string, msg: string): EmbedBuilder {
	const embed = new EmbedBuilder()
		.setColor("Red")
		.setTitle("❌ Failed to start playing")
		.setDescription(msg)
		.addFields([
			{
				name: "Query",
				value: `Tried to fetch: ${query || "None"}`,
				inline: true,
			},
		])
		.setTimestamp();

	return embed;
}
