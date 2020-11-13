package leagueoflegends

import (
	"fmt"

	"github.com/yuhanfang/riot/constants/champion"
)

// Scoring Stores data about score games
type Scoring struct {
	Rank    string
	Winrate string
}

// ProfileLOL Stores data of a profile player
type ProfileLOL struct {
	Scoring
	Champions []champion.Champion
}

// ProfileBuilder Display on Discord profile data
func (profile ProfileLOL) ProfileBuilder() string {
	template := fmt.Sprintf("**%s**\n%s\n\n", profile.Rank, profile.Winrate)
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
