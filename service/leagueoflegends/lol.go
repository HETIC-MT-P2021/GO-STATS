package leagueoflegends

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/wyllisMonteiro/GO-STATS/service/config"
	"github.com/yuhanfang/riot/apiclient"
	"github.com/yuhanfang/riot/constants/champion"
	"github.com/yuhanfang/riot/constants/region"
	"github.com/yuhanfang/riot/ratelimit"
)

const (
	ChampionsLimit int = 3
)

var QueuesType []string = []string{"RANKED_SOLO_5x5"}

func NewConfigLOLAPI(riotGamesToken string) config.ConfigLeagueOfLegendsAPI {
	httpClient := http.DefaultClient
	limiter := ratelimit.NewLimiter()

	return config.ConfigLeagueOfLegendsAPI{
		RiotGamesToken: riotGamesToken,
		Ctx:            context.Background(),
		Limiter:        limiter,
		Client:         apiclient.New(riotGamesToken, httpClient, limiter),
		Region:         region.EUW1,
	}
}

var configAPI config.ConfigLeagueOfLegendsAPI

func GetLOLProfileData(username string) (int, string, string, error) {
	configAPI = NewConfigLOLAPI(os.Getenv("RIOTGAMES"))

	summonerInfos, err := configAPI.Client.GetBySummonerName(configAPI.Ctx, configAPI.Region, username)
	if err != nil {
		return 0, "", "", err
	}

	profileIconID := summonerInfos.ProfileIconID

	summonerChamps, err := GetAllChampionMasteries(summonerInfos.ID)
	if err != nil {
		return profileIconID, "", summonerInfos.Name, err
	}

	rank, winrate, err := GetAllLeaguePositionsForSummoner(summonerInfos.ID)
	if err != nil {
		return profileIconID, "", summonerInfos.Name, err
	}

	profile := ProfileLOL{
		SummonerLevel: summonerInfos.SummonerLevel,
		Rank:          rank,
		Winrate:       winrate,
		Champions:     summonerChamps,
	}

	template := profile.ProfileBuilder()

	prettyPrint(summonerInfos, err)

	return profileIconID, template, summonerInfos.Name, nil
}

func GetAllChampionMasteries(summonerID string) ([]champion.Champion, error) {
	summonerChamps, err := configAPI.Client.GetAllChampionMasteries(configAPI.Ctx, configAPI.Region, summonerID)
	if err != nil {
		return []champion.Champion{}, err
	}

	var filteredChamps []champion.Champion
	for _, champ := range summonerChamps[0:ChampionsLimit] {
		filteredChamps = append(filteredChamps, champ.ChampionID)
	}

	return filteredChamps, nil
}

func GetAllLeaguePositionsForSummoner(SummonerID string) (string, string, error) {
	rank := ""
	winrate := ""

	rankedByModes, err := configAPI.Client.GetAllLeaguePositionsForSummoner(configAPI.Ctx, configAPI.Region, SummonerID)
	if err != nil {
		log.Println(err)
		return "", "", err
	}

	for _, ranked := range rankedByModes {
		found := findQueueType(QueuesType, ranked.QueueType)
		if found {
			rank = fmt.Sprintf("%s %s | %d LP", ranked.Tier, ranked.Rank, ranked.LeaguePoints)
			winrate = fmt.Sprintf("%.2f%% W/L", float64(ranked.Wins)/(float64(ranked.Wins)+float64(ranked.Losses))*100)
		}
	}

	prettyPrint(rankedByModes, err)

	return rank, winrate, nil
}

func findQueueType(fromValues []string, lookingFor string) bool {
	for _, from := range fromValues {
		if from == lookingFor {
			return true
		}
	}
	return false
}

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
