package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/wyllisMonteiro/GO-STATS/service/discord"
	"github.com/wyllisMonteiro/GO-STATS/service/leagueoflegends"
	"github.com/wyllisMonteiro/GO-STATS/service/leagueoflegends/impl"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
)

const (
	// HelpMessage written to help user using command -gs help
	HelpMessage string = "" +
		"- help: displays a list of commands \n" +
		"- stats <game>: displays stats of all players from a specific game \n" +
		"- <player> stats <game>: displays stats of a specific player from a specific game \n"

	defaultDiscordColor int = 0x4E6F7B
	errorDiscordColor   int = 0xA62019
)

// RunCommand Displays data to Discord with given command
func RunCommand(discordBot discord.Bot, messager *discordgo.MessageCreate, args []string) {
	var params = strings.Split(args[1], " ")
	whichCommandToExecute(discordBot, messager, params)
}

func whichCommandToExecute(discordBot discord.Bot, messager *discordgo.MessageCreate, params []string) {
	switch strings.ReplaceAll(params[0], " ", "") {
	// Games here
	case "lol":
		lolCommand(discordBot, messager, params)
	// Bot commands here
	case "help", " ?", "h":
		discordBot.ChannelMessageSendEmbed(messager.ChannelID, embed.NewGenericEmbed("GameStats BOT Helper", HelpMessage))
	case "version", "v":
		discordBot.ChannelMessageSend(messager.ChannelID, "v0.0.1")
	case "clear":
		discordBot.ChannelMessageSend(messager.ChannelID, "Error, please retry...")
	}
}

func lolCommand(discordBot discord.Bot, messager *discordgo.MessageCreate, params []string) {
	if len(params) == 1 {
		discordBot.ChannelMessageSend(messager.ChannelID, "Merci de spécifier le nom du joueur pour League of Legends")
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
	discordBot.ChannelMessageSendEmbed(messager.ChannelID, returnedMessage.MessageEmbed)
}
