package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"net/http"

	"github.com/yuhanfang/riot/apiclient"
	"github.com/yuhanfang/riot/constants/region"
	"github.com/yuhanfang/riot/ratelimit"
)

const (
	reg = region.EUW1
)

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

	//matches := GetMatchList(ctx, reg, summoner.AccountID)

	profileIconID := summoner.ProfileIconID
	data := fmt.Sprintf("- Level %d", summoner.SummonerLevel /*, matches*/)
	summonerName := summoner.Name

	return profileIconID, data, summonerName, nil
}

func GetMatchList(ctx context.Context, r region.Region, accoundID string) string {

	return ""
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
