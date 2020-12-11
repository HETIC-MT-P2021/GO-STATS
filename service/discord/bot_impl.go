package discord

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordBotImpl struct {
	Session *discordgo.Session
}

func (discord DiscordBotImpl) ChannelMessageSendEmbed(a string, b *discordgo.MessageEmbed) {
	discord.Session.ChannelMessageSendEmbed(a, b)
}

func (discord DiscordBotImpl) ChannelMessageSend(a string, b string) {
	discord.Session.ChannelMessageSend(a, b)
}
