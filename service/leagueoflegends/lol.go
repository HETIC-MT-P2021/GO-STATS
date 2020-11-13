package leagueoflegends

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/wyllisMonteiro/GO-STATS/service/config"
	"github.com/yuhanfang/riot/apiclient"
	"github.com/yuhanfang/riot/constants/champion"
	"github.com/yuhanfang/riot/constants/region"
	"github.com/yuhanfang/riot/ratelimit"
)

const (
	championsLimit int = 3
)

// DiscordEmbed Set up NewEmbed .SetThumbnail() || .SetTile() || .Description()
type DiscordEmbed struct {
	ProfileIconID int
	Title         DiscordEmbedTitle
	Description   string
}

// DiscordEmbedTitle Set up NewEmbed.SetTile()
type DiscordEmbedTitle struct {
	SummonerName  string
	SummonerLevel int64
}

var queuesType []string = []string{"RANKED_SOLO_5x5"}

// NewConfigLOLAPI Set up League Of Legends API
func NewConfigLOLAPI(riotGamesToken string) config.LeagueOfLegendsAPI {
	httpClient := http.DefaultClient
	limiter := ratelimit.NewLimiter()

	return config.LeagueOfLegendsAPI{
		RiotGamesToken: riotGamesToken,
		Ctx:            context.Background(),
		Limiter:        limiter,
		Client:         apiclient.New(riotGamesToken, httpClient, limiter),
		Region:         region.EUW1,
	}
}

var configAPI config.LeagueOfLegendsAPI

// GetLOLProfileData Allow to get some data about league of legends player profile from username
func GetLOLProfileData(username string) (DiscordEmbed, error) {
	configAPI = NewConfigLOLAPI(os.Getenv("RIOTGAMES"))

	discordEmbed := DiscordEmbed{}
	summonerInfos, err := configAPI.Client.GetBySummonerName(configAPI.Ctx, configAPI.Region, username)
	if err != nil {
		return discordEmbed, err
	}

	profileIconID := summonerInfos.ProfileIconID

	summonerChamps, err := GetAllChampionMasteries(summonerInfos.ID)
	if err != nil {

		fmt.Println(err)
		return discordEmbed, err
	}

	scoring, err := GetAllLeaguePositionsForSummoner(summonerInfos.ID)
	if err != nil {
		return discordEmbed, err
	}

	profile := ProfileLOL{
		Scoring{
			Rank:    scoring.Rank,
			Winrate: scoring.Winrate,
		},
		summonerChamps,
	}

	template := profile.ProfileBuilder()

	prettyPrint(summonerInfos, err)

	discordEmbed = DiscordEmbed{
		ProfileIconID: profileIconID,
		Title: DiscordEmbedTitle{
			SummonerName:  summonerInfos.Name,
			SummonerLevel: summonerInfos.SummonerLevel,
		},
		Description: template,
	}

	return discordEmbed, nil
}

// GetAllChampionMasteries Allow to get 3 most champion used
func GetAllChampionMasteries(summonerID string) ([]champion.Champion, error) {
	summonerChamps, err := configAPI.Client.GetAllChampionMasteries(configAPI.Ctx, configAPI.Region, summonerID)
	if err != nil {
		return []champion.Champion{}, err
	}

	var filteredChamps []champion.Champion
	if len(summonerChamps) > championsLimit {

		for _, champ := range summonerChamps[0:championsLimit] {
			filteredChamps = append(filteredChamps, champ.ChampionID)
		}
	} else {

		return nil, err
	}

	return filteredChamps, nil
}

// GetAllLeaguePositionsForSummoner Allow to get some data in ranked mode
func GetAllLeaguePositionsForSummoner(SummonerID string) (Scoring, error) {
	var scoring Scoring

	rankedByModes, err := configAPI.Client.GetAllLeaguePositionsForSummoner(configAPI.Ctx, configAPI.Region, SummonerID)
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

			scoring = Scoring{
				Rank:    rank,
				Winrate: winrate,
			}
		}
	}

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

// prettyPrint Made data readable
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

func upFirstCaseLetter(string string) string {
	return strings.Title(strings.ToLower(string))
}
