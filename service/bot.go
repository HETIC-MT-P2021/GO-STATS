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
const CommandPrefix = "-gs"

// DG : Create new session of discord
var discordSession *discordgo.Session

// ConnectBot : Connect bot to server
func ConnectBot() {
	botConfig, err := GetVarsBot()
	if err != nil {
		log.Println(err)
		return
	}

	discordSession, err = discordgo.New("Bot " + botConfig.Token)
	if err != nil {
		log.Println(err)
		return
	}
}

// RunBot : Create new bot
func RunBot() {
	ConnectBot()

	user, err := discordSession.User("@me")
	if err != nil {
		log.Println(err)
		return
	}

	botID = user.ID

	discordSession.AddHandler(MessageHandler)

	err = discordSession.Open()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Bot is running")

	<-make(chan struct{})
	return
}

// MessageHandler : Waiting for sending message by user
func MessageHandler(Session *discordgo.Session, Messager *discordgo.MessageCreate) {
	if Messager.Author.ID == botID {
		return
	}
	var args = strings.Split(Messager.Content, CommandPrefix)

	fmt.Println(args[0])
	runCommands(Session, Messager, args)
}
