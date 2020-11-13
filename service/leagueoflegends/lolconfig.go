package leagueoflegends

import "github.com/wyllisMonteiro/GO-STATS/service/config"

type LeagueOfLegends interface {
	MakeConfig(string) config.ConfigLeagueOfLegendsAPI
}
