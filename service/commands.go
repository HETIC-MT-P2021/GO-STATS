package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/wyllisMonteiro/GO-STATS/service/leagueoflegends"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	"github.com/wyllisMonteiro/GO-STATS/service/leagueoflegends/impl"
)

const (
	helpMessage string = "" +
		"- help: displays a list of commands \n" +
		"- stats <game>: displays stats of all players from a specific game \n" +
		"- <player> stats <game>: displays stats of a specific player from a specific game \n"

	defaultDiscordColor int = 0x4E6F7B
	errorDiscordColor   int = 0xA62019
)

// runCommands Check which command is called
func runCommands(Session *discordgo.Session, Messager *discordgo.MessageCreate, args []string) {

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
		fmt.Println(params)
		if len(params) == 1 {
			Session.ChannelMessageSend(Messager.ChannelID, "Merci de spécifier le nom du joueur pour League of Legends")
			return
		}

		config := leagueoflegends.MakeConfig(os.Getenv("RIOTGAMES"))
		api := &impl.Impl{
			Config: config,
		}

		discordEmbed, err := api.GetLOLProfileData(params[1])

		returnedMessage := embed.NewEmbed()
		if err != nil {
			if err.Error() == "forbidden" {
				returnedMessage.SetTitle("Access Forbidden")
				returnedMessage.SetDescription("Please have a look to API Key before retrying.")
				returnedMessage.SetColor(errorDiscordColor)
			} else {
				returnedMessage.SetTitle("Summoner not found")
				returnedMessage.SetDescription(fmt.Sprintf("No summoner found for username : '%s'", params[1]))
				returnedMessage.SetColor(errorDiscordColor)
			}
		} else {
			returnedMessage.SetThumbnail(fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/10.22.1/img/profileicon/%d.png", discordEmbed.ProfileIconID))
			returnedMessage.SetTitle(fmt.Sprintf("%s - Level %d", discordEmbed.Title.SummonerName, discordEmbed.Title.SummonerLevel))
			returnedMessage.SetDescription(discordEmbed.Description)
			returnedMessage.SetColor(defaultDiscordColor)
		}

		Session.ChannelMessageSendEmbed(Messager.ChannelID, returnedMessage.MessageEmbed)
		//Session.ChannelMessageSend(Messager.ChannelID, "Error, please retry...")

	case "clear":
		Session.ChannelMessageSend(Messager.ChannelID, "Error, please retry...")
	}
}
