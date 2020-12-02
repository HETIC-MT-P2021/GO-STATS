package leagueoflegends

import (
	"github.com/wyllisMonteiro/GO-STATS/service/leagueoflegends/structs"
	"github.com/wyllisMonteiro/GO-STATS/service/templates"
	"github.com/yuhanfang/riot/apiclient"
)

// LeagueOfLegends Stores all methods to use riot API for League Of Legends
type LeagueOfLegends interface {
	GetLOLProfileData(username string) (structs.DiscordEmbed, error)
	GetAllChampionMasteries(summonerID string) ([]apiclient.ChampionMastery, error)
	GetAllLeaguePositionsForSummoner(summonerID string) (templates.Scoring, error)
}
