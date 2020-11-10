package service

import (
	"fmt"
	"github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	"strings"
)

const helpMessage = "" +
	"- help: displays a list of commands \n" +
	"- stats <game>: displays stats of all players from a specific game \n" +
	"- <player> stats <game>: displays stats of a specific player from a specific game \n"

// RunBot : Create new bot
func runCommands(Session *discordgo.Session, Messager *discordgo.MessageCreate, args [] string) {

	var params = strings.Split(args[1], " ")
	fmt.Println(params)

	// for each command, do something different
	switch strings.ReplaceAll(params[0], " ", "") {
	case "help", " ?", "h":
		Session.ChannelMessageSendEmbed(Messager.ChannelID, embed.NewGenericEmbed("GameStats BOT Helper", helpMessage))
	case "version", "v":
		Session.ChannelMessageSend(Messager.ChannelID, "v0.0.1")
	case "me":
		Session.ChannelMessageSend(Messager.ChannelID, Messager.Author.Username)
	case "lol":
		profileIconID, data, summonerName, err := GetLolData(params[1])

		returnedMessage := embed.NewEmbed()
		if err != nil {

			returnedMessage.SetTitle("Summoner not found")
			returnedMessage.SetDescription(fmt.Sprintf("No summoner found for username : '%s'", params[1]))
			returnedMessage.SetColor(0xA62019)
		} else {

			returnedMessage.SetThumbnail(fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/10.22.1/img/profileicon/%d.png", profileIconID))
			returnedMessage.SetTitle(summonerName)
			returnedMessage.SetDescription(data)
			returnedMessage.SetColor(0x4E6F7B)
		}
		Session.ChannelMessageSendEmbed(Messager.ChannelID, returnedMessage.MessageEmbed)

	case "clear":
		Session.ChannelMessageSend(Messager.ChannelID, "Error, please retry...")
	}
}
