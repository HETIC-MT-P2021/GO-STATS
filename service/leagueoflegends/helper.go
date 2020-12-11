package leagueoflegends

import (
	"context"
	"net/http"

	"github.com/wyllisMonteiro/GO-STATS/service/leagueoflegends/structs"
	"github.com/yuhanfang/riot/apiclient"
	"github.com/yuhanfang/riot/constants/region"
	"github.com/yuhanfang/riot/ratelimit"
)

// MakeConfig Allows to get stuff to make API call
func MakeConfig(riotGamesToken string) structs.LeagueOfLegendsAPI {
	httpClient := http.DefaultClient
	limiter := ratelimit.NewLimiter()

	return structs.LeagueOfLegendsAPI{
		RiotGamesToken: riotGamesToken,
		Ctx:            context.Background(),
		Limiter:        limiter,
		Client:         apiclient.New(riotGamesToken, httpClient, limiter),
		Region:         region.EUW1,
	}
}
