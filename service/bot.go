package service

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"strings"
)

// BotConfig : config to create new bot
type BotConfig struct {
	Token string
}

var botID string
var PREFIX = "-gs"

// DG : Create new session of discord
var DG *discordgo.Session

// ConnectBot : Connect bot to server
func ConnectBot() {
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
	ConnectBot()

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
	var args = strings.Split(m.Content, PREFIX)

	fmt.Println(args[0])
	runCommands(s, m, args)
}
