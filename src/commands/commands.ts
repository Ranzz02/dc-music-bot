import {
	ApplicationCommandOptionType,
	RESTPostAPIChatInputApplicationCommandsJSONBody,
} from "discord.js";
import { Commands } from "../core/types";

export type CommandDefinition = RESTPostAPIChatInputApplicationCommandsJSONBody;

export const PlayCommand: CommandDefinition = {
	name: Commands.Play,
	description: "Play a song from youtube",
	options: [
		{
			name: "query",
			description: "URL or query string",
			type: ApplicationCommandOptionType.String,
			required: true,
		},
	],
};

export const PingCommand: CommandDefinition = {
	name: "ping",
	description: "Responds with Pong!",
};
