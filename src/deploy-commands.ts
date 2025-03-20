import { REST, Routes, SlashCommandBuilder } from "discord.js";
import dotenv from "dotenv";
import { PingCommand, PlayCommand } from "./commands/commands";

dotenv.config();

const commands = [PingCommand, PlayCommand];

const rest = new REST({ version: "10" }).setToken(process.env.DISCORD_TOKEN!);

export async function registerCommands() {
	try {
		console.log("Registering slash commands...");
		await rest.put(Routes.applicationCommands(process.env.CLIENT_ID!), {
			body: commands,
		});
		console.log("Slash commands registered successfully!");
	} catch (error) {
		console.error(error);
	}
}
