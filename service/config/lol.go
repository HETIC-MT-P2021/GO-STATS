package config

import (
	"context"

	"github.com/yuhanfang/riot/apiclient"
	"github.com/yuhanfang/riot/constants/region"
	"github.com/yuhanfang/riot/ratelimit"
)

type ConfigLeagueOfLegendsAPI struct {
	RiotGamesToken string
	Ctx            context.Context
	Limiter        ratelimit.Limiter
	Client         apiclient.Client
	Region         region.Region
}
