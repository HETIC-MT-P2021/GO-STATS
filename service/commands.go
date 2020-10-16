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

	switch strings.ReplaceAll(args[1], " ", "") {
	case "help":
		Session.ChannelMessageSendEmbed(Messager.ChannelID, embed.NewGenericEmbed("GameStats BOT Helper", helpMessage))
	case "version":
		Session.ChannelMessageSend(Messager.ChannelID, "v0.0.1")
	case "lol":
		Session.ChannelMessageSend(Messager.ChannelID, "League of Legends")
	case "clear":
		Session.ChannelMessageSend(Messager.ChannelID, "Error, please retry...")
	}
}
