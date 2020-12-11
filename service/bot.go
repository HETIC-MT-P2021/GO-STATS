package service

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/wyllisMonteiro/GO-STATS/service/command"
	"github.com/wyllisMonteiro/GO-STATS/service/discord"
)

// BotConfig Stores necessary data to create new bot
type BotConfig struct {
	Token string
}

var botID string

// discordSession Create new session of discord
var discordSession *discordgo.Session

// ConnectBot Make bot used on a server
func ConnectBot() {
	botConfig, err := GetVarsBot()
	if err != nil {
		log.Println(err)
		fmt.Println(-2)
		return
	}

	discordSession, err = discordgo.New("Bot " + botConfig.Token)
	if err != nil {
		log.Println(err)
		fmt.Println(-1)
		return
	}
}

// RunBot Make all features usable
func RunBot() {
	ConnectBot()

	user, err := discordSession.User("@me")
	if err != nil {
		log.Println(err)
		fmt.Println(0)
		return
	}

	botID = user.ID

	discordSession.AddHandler(MessageHandler)

	err = discordSession.Open()
	if err != nil {
		log.Println(err)
		fmt.Println(1)
		return
	}

	fmt.Println("Bot is running")

	<-make(chan struct{})
	return
}

// MessageHandler Waiting for sending message by user
func MessageHandler(session *discordgo.Session, messager *discordgo.MessageCreate) {
	if messager.Author.ID == botID {
		return
	}
	var args = strings.Split(messager.Content, os.Getenv("COMMANDPREFIX"))

	discord := discord.DiscordBotImpl{
		Session: session,
	}

	command.RunCommand(discord, messager, args)
}
