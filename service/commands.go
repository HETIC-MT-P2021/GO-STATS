package service

import (
	"github.com/bwmarrin/discordgo"

	"strings"
)

// RunBot : Create new bot
func runCommands(s *discordgo.Session, m *discordgo.MessageCreate, args [] string) {

	switch strings.ReplaceAll(args[1], " ", "") {
	case "ping":
		s.ChannelMessageSend(m.ChannelID, "pong")
	case "version":
		s.ChannelMessageSend(m.ChannelID, "v0.0.1")
	case "lol":
		s.ChannelMessageSend(m.ChannelID, "League of Legends")
	case "clear":
		s.ChannelMessageSend(m.ChannelID, "Error, please retry...")
	}
}
