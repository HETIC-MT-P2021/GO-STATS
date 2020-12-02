package leagueoflegends

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/wyllisMonteiro/GO-STATS/service/leagueoflegends/structs"
	"github.com/wyllisMonteiro/GO-STATS/service/mocks"
	"github.com/wyllisMonteiro/GO-STATS/service/templates"
	"github.com/yuhanfang/riot/apiclient"
)

func TestGetLOLProfileData(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mock := mocks.NewMockLeagueOfLegends(controller)

	username := "magma"
	mock.EXPECT().GetLOLProfileData(username).Return(structs.DiscordEmbed{}, nil)
	mock.GetLOLProfileData(username)
}

func TestGetAllChampionMasteries(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mock := mocks.NewMockLeagueOfLegends(controller)

	summonerID := "1"
	mock.EXPECT().GetAllChampionMasteries(summonerID).Return([]apiclient.ChampionMastery{}, nil)
	mock.GetAllChampionMasteries(summonerID)
}

func TestGetAllLeaguePositionsForSummoner(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mock := mocks.NewMockLeagueOfLegends(controller)

	summonerID := "1"
	mock.EXPECT().GetAllLeaguePositionsForSummoner(summonerID).Return(templates.Scoring{}, nil)
	mock.GetAllLeaguePositionsForSummoner(summonerID)
}
