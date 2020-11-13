package leagueoflegends

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/wyllisMonteiro/GO-STATS/service/config"
	"github.com/wyllisMonteiro/GO-STATS/service/mocks"
	"github.com/yuhanfang/riot/apiclient"
	"github.com/yuhanfang/riot/constants/region"
	"github.com/yuhanfang/riot/ratelimit"
)

func fakeConfigLeagueOfLegendsAPI(riotGamesToken string) config.ConfigLeagueOfLegendsAPI {
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

func TestLeagueOfLegends(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mock := mocks.NewMockLeagueOfLegends(controller)

	wantedConfig := NewConfigLOLAPI("token")

	mock.EXPECT().MakeConfig("token").Return(wantedConfig)
	mock.MakeConfig("token")
}
