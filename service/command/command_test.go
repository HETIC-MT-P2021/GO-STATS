package command

import (
	"fmt"
	"os"
	"strings"
	"testing"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	"github.com/golang/mock/gomock"
	"github.com/wyllisMonteiro/GO-STATS/service/leagueoflegends"
	"github.com/wyllisMonteiro/GO-STATS/service/leagueoflegends/impl"
	"github.com/wyllisMonteiro/GO-STATS/service/leagueoflegends/structs"
	"github.com/wyllisMonteiro/GO-STATS/service/mocks"
)

func createFakeMessager() *discordgo.MessageCreate {
	return &discordgo.MessageCreate{
		&discordgo.Message{
			ID:        "10",
			ChannelID: "1",
			GuildID:   "",
			Content:   "-gs lol Xari", // Here is command
		},
	}
}

func TestRunCommandLol(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	returnedMessage := embed.NewEmbed()
	messager := createFakeMessager()
	session := mocks.NewMockBot(controller)

	session.EXPECT().ChannelMessageSendEmbed(messager.ChannelID, returnedMessage.MessageEmbed)

	args := [2]string{" ", "lol Xari"}

	var params = strings.Split(args[1], " ")

	if params[0] != "lol" {
		t.Errorf("current value " + params[0] + " wanted value : lol")
	}

	whichCommandToExecute(session, messager, params)
}

func TestRunCommandHelp(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	messager := createFakeMessager()
	session := mocks.NewMockBot(controller)
	session.EXPECT().ChannelMessageSendEmbed(messager.ChannelID,
		embed.NewGenericEmbed("GameStats BOT Helper",
			HelpMessage))

	args := [2]string{" ", "help"}

	var params = strings.Split(args[1], " ")

	if params[0] != "help" {
		t.Errorf("current value " + params[0] + " wanted value : help")
	}

	whichCommandToExecute(session, messager, params)
}

func TestRunCommandVersion(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	messager := createFakeMessager()
	session := mocks.NewMockBot(controller)
	session.EXPECT().ChannelMessageSend(messager.ChannelID, "v0.0.1")

	args := [2]string{" ", "version"}

	var params = strings.Split(args[1], " ")

	if params[0] != "version" {
		t.Errorf("current value " + params[0] + " wanted value : version")
	}

	whichCommandToExecute(session, messager, params)
}

func TestRunCommandClear(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	messager := createFakeMessager()
	session := mocks.NewMockBot(controller)
	session.EXPECT().ChannelMessageSend(messager.ChannelID, "Error, please retry...")

	args := [2]string{" ", "clear"}

	var params = strings.Split(args[1], " ")

	if params[0] != "clear" {
		t.Errorf("current value " + params[0] + " wanted value : clear")
	}

	whichCommandToExecute(session, messager, params)
}

func TestLolCommandLol_missing_params(t *testing.T) {
	params := [1]string{"lol"}
	if len(params) == 1 {
		controller := gomock.NewController(t)
		defer controller.Finish()

		messager := createFakeMessager()
		session := mocks.NewMockBot(controller)
		session.EXPECT().ChannelMessageSend(messager.ChannelID, "Merci de spécifier le nom du joueur pour League of Legends")

		session.ChannelMessageSend(messager.ChannelID, "Merci de spécifier le nom du joueur pour League of Legends")
	} else {
		t.Errorf(fmt.Sprintf("current value : %d wanted value : 1", len(params)))
	}
}

func TestLolCommandLol(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	messager := createFakeMessager()
	session := mocks.NewMockBot(controller)

	params := [2]string{"lol", "Xari"}

	if len(params) == 1 {
		session.EXPECT().ChannelMessageSend(messager.ChannelID, "Merci de spécifier le nom du joueur pour League of Legends")
		session.ChannelMessageSend(messager.ChannelID, "Merci de spécifier le nom du joueur pour League of Legends")
		return
	}

	config := leagueoflegends.MakeConfig(os.Getenv("RIOTGAMES"))
	api := &impl.Impl{
		Config: config,
	}

	leagueoflegendsMock := mocks.NewMockLeagueOfLegends(controller)

	username := "magma"
	leagueoflegendsMock.EXPECT().GetLOLProfileData(username).Return(structs.DiscordEmbed{}, nil)
	leagueoflegendsMock.GetLOLProfileData(username)

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

	session.EXPECT().ChannelMessageSendEmbed(messager.ChannelID, returnedMessage.MessageEmbed)
	session.ChannelMessageSendEmbed(messager.ChannelID, returnedMessage.MessageEmbed)
}
