package service

import (
	"fmt"

	"github.com/yuhanfang/riot/constants/champion"
)

type ProfileLOL struct {
	SummonerLevel int64
	Rank          string
	Winrate       string
	Champions     []champion.Champion
}

func (profile ProfileLOL) ProfileBuilder() string {
	template := fmt.Sprintf("\n- **Level %d**\n", profile.SummonerLevel)
	template += fmt.Sprintf("%s\n%s\n", profile.Rank, profile.Winrate)
	template += "- **Champions : **"

	championsLength := len(profile.Champions) - 1
	for index, championItem := range profile.Champions {
		if index == championsLength {
			template += fmt.Sprintf("%s", championItem)
		} else {
			template += fmt.Sprintf("%s, ", championItem)
		}
	}

	return template
}
