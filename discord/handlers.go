package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandsHandlers Bot command handler
func CommandsHandlers(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore any messages sent by the bot
	if message.Author.ID == session.State.User.ID {
		return
	}

	// Ignore any message that doesn't have the prefix !untitled
	if !strings.HasPrefix(message.Content, "!untitled") {
		return
	}

	// Grab command args (ignore prefix)
	cmdArgs := strings.Split(message.Content, " ")[1:]
	// Check for commands present
	if len(cmdArgs) == 0 {
		session.ChannelMessageSend(message.ChannelID, errorMessage("Missing command!", "For a help typed !untitled help"))
		return
	}

	switch cmdArgs[0] {
	case "ping":
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	case "pong":
		session.ChannelMessageSend(message.ChannelID, "Ping!")
	case "plagueis":
		session.ChannelMessageSend(message.ChannelID, "Did you ever hear the Tragedy of Darth Plagueis the wise? I thought not. It's not a story the Jedi would tell you. It's a Sith legend. Darth Plagueis was a Dark Lord of the Sith, so powerful and so wise he could use the Force to influence the midichlorians to create life... He had such a knowledge of the dark side that he could even keep the ones he cared about from dying. The dark side of the Force is a pathway to many abilities some consider to be unnatural. He became so powerful... the only thing he was afraid of was losing his power, which eventually, of course, he did. Unfortunately, he taught his apprentice everything he knew, then his apprentice killed him in his sleep. It's ironic he could save others from death, but not himself.")
	case "hello":
		session.ChannelMessageSend(message.ChannelID, "World!")
	case "help":
		if len(cmdArgs) > 1 {
			helpCommandHandler(session, message, cmdArgs[1])
		} else {
			helpCommandHandler(session, message, "")
		}
	default:
		session.ChannelMessageSend(message.ChannelID, errorMessage("Invalid command!", "For a list of commands/help, type `!unititled help`"))
	}
}

// Generic message format for errors
func errorMessage(title string, message string) string {
	return "❌  **" + title + "**\n" + message
}
