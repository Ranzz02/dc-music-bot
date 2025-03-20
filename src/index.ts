import { Client, GatewayIntentBits, Interaction } from "discord.js";
import dotenv from "dotenv";
import { registerCommands } from "./deploy-commands";
import { Commands } from "./core/types";
import { Play } from "./handlers/music";
import ffmpeg from "ffmpeg-static";
process.env.FFMPEG_PATH = ffmpeg as string;

dotenv.config();

const client = new Client({
	intents: [
		GatewayIntentBits.Guilds,
		GatewayIntentBits.GuildMessages,
		GatewayIntentBits.MessageContent,
	],
});

client.once("ready", async (client) => {
	console.log(`Logged in as ${client.user?.tag}!`);
	await registerCommands();
});

// Slash command Interactions
client.on("interactionCreate", async (interaction: Interaction) => {
	if (!interaction.isChatInputCommand()) return;

	if (interaction.commandName === Commands.Play) {
		Play(interaction);
	}
});

client.login(process.env.DISCORD_TOKEN);
