package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yuhanfang/riot/apiclient"
	"github.com/yuhanfang/riot/constants/region"
	"github.com/yuhanfang/riot/constants/champion"
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

	summonerInfos, err := client.GetBySummonerName(ctx, reg, username)
	if err != nil {

		return 0, "", "", err
	}
	prettyPrint(summonerInfos, err)

	summonerChamps, err := client.GetAllChampionMasteries(ctx, reg, summonerInfos.ID)
	if err != nil {

		return 0, "", "", err
	}

	var filteredChamps []champion.Champion
	for _, champ := range summonerChamps[0:3] {

		filteredChamps = append(filteredChamps, champ.ChampionID)
	}

	profileIconID := summonerInfos.ProfileIconID
	data := fmt.Sprintf("\n- **Level %d**\n- **Champions : **%s, %s, %s", summonerInfos.SummonerLevel, filteredChamps[0], filteredChamps[1], filteredChamps[2])
	summonerName := summonerInfos.Name

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