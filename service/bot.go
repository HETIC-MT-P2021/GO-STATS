package service

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

// BotConfig : config to create new bot
type BotConfig struct {
	Token string
}

var botID string
const PREFIX = '-'

// DG : Create new session of discord
var DG *discordgo.Session

// CnnectBot : Connect bot to server
func CnnectBot() {
	botConfig, err := GetVarsBot()
	if err != nil {
		log.Println(err)
		return
	}

	DG, err = discordgo.New("Bot " + botConfig.Token)
	if err != nil {
		log.Println(err)
		return
	}
}

// RunBot : Create new bot
func RunBot() {
	CnnectBot()

	u, err := DG.User("@me")
	if err != nil {
		log.Println(err)
		return
	}

	botID = u.ID

	DG.AddHandler(MessageHandler)

	err = DG.Open()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Bot is running")

	<-make(chan struct{})
	return
}

// MessageHandler : Waiting for sending message by user
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == botID {
		return
	}

	var args = strings.Split(m.Content, "-")
	fmt.Println(args)
	if m.Content == "-gs stats lol" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Soon")
	}
}
