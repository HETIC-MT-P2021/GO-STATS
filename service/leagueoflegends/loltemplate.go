package leagueoflegends

import (
	"fmt"
	"github.com/yuhanfang/riot/apiclient"
)

// Scoring Stores data about score games
type Scoring struct {
	Rank    string
	Winrate string
}

// ProfileLOL Stores data of a profile player
type ProfileLOL struct {
	Scoring
	Champions []apiclient.ChampionMastery
}

func (profile ProfileLOL) ProfileBuilder() string {

	template := ""
	if len(profile.Rank) == 0 {
		template += "This player is not rank in Solo/Duo\n"
	} else {
		template += fmt.Sprintf("**%s**\n%s\n", profile.Rank, profile.Winrate)
	}
	template += "\n- **Champions : **"
	template += fmt.Sprintf("\n\n > :one: %s - %d pts", profile.Champions[0].ChampionID, profile.Champions[0].ChampionPoints)
	template += fmt.Sprintf("\n > :two: %s - %d pts", profile.Champions[1].ChampionID, profile.Champions[1].ChampionPoints)
	template += fmt.Sprintf("\n > :three: %s - %d pts", profile.Champions[2].ChampionID, profile.Champions[2].ChampionPoints)

	return template
}
