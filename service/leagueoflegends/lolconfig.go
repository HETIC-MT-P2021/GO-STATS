package leagueoflegends

import "github.com/wyllisMonteiro/GO-STATS/service/config"

// LeagueOfLegends Set up league of legends config
type LeagueOfLegends interface {
	MakeConfig(string) config.LeagueOfLegendsAPI
}
