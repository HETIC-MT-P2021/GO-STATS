package leagueoflegends

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/wyllisMonteiro/GO-STATS/service/mocks"
)

func TestLeagueOfLegends(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mock := mocks.NewMockLeagueOfLegends(controller)

	wantedConfig := NewConfigLOLAPI("token")

	mock.EXPECT().MakeConfig("token").Return(wantedConfig)
	mock.MakeConfig("token")
}
