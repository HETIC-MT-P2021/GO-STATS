package impl

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/wyllisMonteiro/GO-STATS/service/leagueoflegends/constants"
	"github.com/wyllisMonteiro/GO-STATS/service/leagueoflegends/structs"
	"github.com/wyllisMonteiro/GO-STATS/service/templates"
	"github.com/yuhanfang/riot/apiclient"
)

var queuesType []string = []string{"RANKED_SOLO_5x5"}

// Impl Implementation of LeagueOfLegends interface
type Impl struct {
	Config structs.LeagueOfLegendsAPI
}

// GetLOLProfileData Allow to get some data about league of legends player profile from username
func (leagueOfLegends Impl) GetLOLProfileData(username string) (structs.DiscordEmbed, error) {
	discordEmbed := structs.DiscordEmbed{}

	config := leagueOfLegends.Config

	summonerInfos, err := config.Client.GetBySummonerName(config.Ctx, config.Region, username)
	if err != nil {
		return discordEmbed, err
	}

	profileIconID := summonerInfos.ProfileIconID

	summonerChamps, err := leagueOfLegends.GetAllChampionMasteries(summonerInfos.ID)
	if err != nil {
		fmt.Println(err)
		return discordEmbed, err
	}

	scoring, err := leagueOfLegends.GetAllLeaguePositionsForSummoner(summonerInfos.ID)
	if err != nil {
		return discordEmbed, err
	}

	profile := templates.ProfileLOL{
		templates.Scoring{
			Rank:    scoring.Rank,
			Winrate: scoring.Winrate,
		},
		summonerChamps,
	}

	template := profile.ProfileBuilder()

	prettyPrint(summonerInfos, err)

	discordEmbed = structs.DiscordEmbed{
		ProfileIconID: profileIconID,
		Title: structs.DiscordEmbedTitle{
			SummonerName:  summonerInfos.Name,
			SummonerLevel: summonerInfos.SummonerLevel,
		},
		Description: template,
	}

	return discordEmbed, nil
}

// GetAllChampionMasteries Allow to get 3 most champion used
func (leagueOfLegends Impl) GetAllChampionMasteries(summonerID string) ([]apiclient.ChampionMastery, error) {
	config := leagueOfLegends.Config

	summonerChamps, err := config.Client.GetAllChampionMasteries(config.Ctx, config.Region, summonerID)
	if err != nil {
		return []apiclient.ChampionMastery{}, err
	}

	var filteredChamps []apiclient.ChampionMastery
	if len(summonerChamps) > constants.ChampionsLimit {
		for _, champ := range summonerChamps[0:constants.ChampionsLimit] {
			filteredChamps = append(filteredChamps, champ)
		}
	} else {

		return nil, err
	}

	return filteredChamps, nil
}

// GetAllLeaguePositionsForSummoner Allow to get some data in ranked mode
func (leagueOfLegends Impl) GetAllLeaguePositionsForSummoner(SummonerID string) (templates.Scoring, error) {
	config := leagueOfLegends.Config

	var scoring templates.Scoring

	rankedByModes, err := config.Client.GetAllLeaguePositionsForSummoner(config.Ctx, config.Region, SummonerID)
	if err != nil {
		log.Println(err)
		return scoring, err
	}

	prettyPrint(rankedByModes, err)

	for _, ranked := range rankedByModes {
		found := findRankedSoloDuo(queuesType, ranked.QueueType)
		if found {
			rank := fmt.Sprintf("%s %s\n > %d LP", upFirstCaseLetter(string(ranked.Tier)), ranked.Rank, ranked.LeaguePoints)
			winrate := fmt.Sprintf("%.2f%% W/L", float64(ranked.Wins)/(float64(ranked.Wins)+float64(ranked.Losses))*100)

			scoring = templates.Scoring{
				Rank:    rank,
				Winrate: winrate,
			}
		}
	}

	prettyPrint(rankedByModes, err)
	return scoring, nil
}

// findRankedSoloDuo Allow to get some data in ranked mode
func findRankedSoloDuo(fromValues []string, lookingFor string) bool {
	for _, from := range fromValues {
		if from == lookingFor {
			return true
		}
	}
	return false
}

// prettyPrint Makes data readable
func prettyPrint(res interface{}, err error) {
	if err != nil {
		fmt.Println("HTTP error:", err)
		return
	}
	js, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}
	fmt.Println(string(js))
}

// upFirstCaseLetter
func upFirstCaseLetter(string string) string {
	return strings.Title(strings.ToLower(string))
}
