package templates

import (
	"fmt"

	"github.com/wyllisMonteiro/GO-STATS/service/leagueoflegends/constants"
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

// ProfileBuilder Displays on discord player profile
func (profile ProfileLOL) ProfileBuilder() string {
	template := ""
	if profile.Rank == "" {
		template += "Not ranked in Solo/Duo\n\n"
	} else {
		template += fmt.Sprintf("**%s**\n%s\n\n", profile.Rank, profile.Winrate)
	}

	if len(profile.Champions) == constants.ChampionsLimit {
		template += "- **Champions : **\n\n"
		template += fmt.Sprintf(" > :one: %s - %d pts\n", profile.Champions[0].ChampionID, profile.Champions[0].ChampionPoints)
		template += fmt.Sprintf(" > :two: %s - %d pts\n", profile.Champions[1].ChampionID, profile.Champions[1].ChampionPoints)
		template += fmt.Sprintf(" > :three: %s - %d pts", profile.Champions[2].ChampionID, profile.Champions[2].ChampionPoints)
	} else {
		template += "- **Champions : **Not enough champions to display"
	}

	return template
}
