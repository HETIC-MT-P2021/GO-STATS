package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/yuhanfang/riot/apiclient"
	"github.com/yuhanfang/riot/constants/region"
	"github.com/yuhanfang/riot/ratelimit"
)

const (
	MatchLimit int = 3
)

var QueuesType []string = []string{"RANKED_SOLO_5x5"}

type ConfigLolAPI struct {
	RiotGamesToken string
	Ctx            context.Context
	Limiter        ratelimit.Limiter
	Client         apiclient.Client
	Region         region.Region
}

var configAPI ConfigLolAPI

func NewConfigLolAPI(riotGamesToken string, reg string) ConfigLolAPI {
	httpClient := http.DefaultClient
	limiter := ratelimit.NewLimiter()

	return ConfigLolAPI{
		RiotGamesToken: riotGamesToken,
		Ctx:            context.Background(),
		Limiter:        limiter,
		Client:         apiclient.New(riotGamesToken, httpClient, limiter),
		Region:         region.EUW1,
	}
}

func GetLolData(username string) (int, string, string, error) {
	configAPI = NewConfigLolAPI(os.Getenv("RIOTGAMES"), region.EUW1)

	fmt.Println("GetBySummonerName")
	summoner, err := configAPI.Client.GetBySummonerName(configAPI.Ctx, configAPI.Region, username)
	if err != nil {
		return 0, "", "", err
	}
	prettyPrint(summoner, err)

	profileIconID := summoner.ProfileIconID
	data := fmt.Sprintf("- Level %d\n", summoner.SummonerLevel)

	rankedSlice, err := GetAllLeaguePositionsForSummoner(summoner.ID)
	for _, ranked := range rankedSlice {
		data += fmt.Sprintf("%s\n", ranked)
	}

	return profileIconID, data, summoner.Name, nil
}

func GetAllLeaguePositionsForSummoner(SummonerID string) ([]string, error) {
	var rankedDatas []string

	rankedByModes, err := configAPI.Client.GetAllLeaguePositionsForSummoner(configAPI.Ctx, configAPI.Region, SummonerID)
	if err != nil {
		log.Println(err)
		return rankedDatas, err
	}

	for _, ranked := range rankedByModes {
		found := findQueueType(QueuesType, ranked.QueueType)
		if found {
			rank := fmt.Sprintf("%s %s | %d LP", ranked.Tier, ranked.Rank, ranked.LeaguePoints)
			winrate := fmt.Sprintf("%.2f%% W/L", float64(ranked.Wins)/(float64(ranked.Wins)+float64(ranked.Losses))*100)
			rankedDatas = append(rankedDatas, rank)
			rankedDatas = append(rankedDatas, winrate)
		}
	}

	prettyPrint(rankedByModes, err)
	return rankedDatas, nil
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
