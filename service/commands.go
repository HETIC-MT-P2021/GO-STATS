package service

import (
	"github.com/bwmarrin/discordgo"
	"github.com/Clinet/discordgo-embed"
	"strings"
)

const helpMessage = "" +
	"- help: displays a list of commands \n" +
	"- stats <game>: displays stats of all players from a specific game \n" +
	"- <player> stats <game>: displays stats of a specific player from a specific game \n"

// RunBot : Create new bot
func runCommands(Session *discordgo.Session, Messager *discordgo.MessageCreate, args [] string) {

	// for each command, do something different
	switch strings.ReplaceAll(args[1], " ", "") {
	case "help", "?", "h":
		Session.ChannelMessageSendEmbed(Messager.ChannelID, embed.NewGenericEmbed("GameStats BOT Helper", helpMessage))
	case "version":
		Session.ChannelMessageSend(Messager.ChannelID, "v0.0.1")
	case "me":
		Session.ChannelMessageSend(Messager.ChannelID, Messager.Author.Username)
	case "lol":
		Session.ChannelMessageSend(Messager.ChannelID, getLolData(Messager.Author.Username))
	case "clear":
		Session.ChannelMessageSend(Messager.ChannelID, "Error, please retry...")
	}
}
