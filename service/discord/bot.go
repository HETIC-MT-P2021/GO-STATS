package discord

import "github.com/bwmarrin/discordgo"

// Bot Stores all methods to display message in Discord
type Bot interface {
	ChannelMessageSendEmbed(string, *discordgo.MessageEmbed)
	ChannelMessageSend(string, string)
}
