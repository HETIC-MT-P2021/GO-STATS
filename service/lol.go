package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yuhanfang/riot/apiclient"
	"github.com/yuhanfang/riot/constants/region"
	"github.com/yuhanfang/riot/ratelimit"
	"net/http"
	"os"
)

const (
	reg = region.EUW1
)

type MessageEmbedImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

func GetLolData(username string) (int, string, string, error) {

	key := os.Getenv("RIOTGAMES")
	httpClient := http.DefaultClient
	ctx := context.Background()
	limiter := ratelimit.NewLimiter()
	client := apiclient.New(key, httpClient, limiter)

	fmt.Println("GetBySummonerName")
	summoner, err := client.GetBySummonerName(ctx, reg, username)
	if err != nil {

		return 0, "", "", err
	}
	prettyPrint(summoner, err)

	profileIconID := summoner.ProfileIconID
	data := fmt.Sprintf("- Level %d\n- Champions : %s, %s, %s", summoner.SummonerLevel, )
	summonerName := summoner.Name

	return profileIconID, data, summonerName, nil
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