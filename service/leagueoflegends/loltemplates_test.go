package leagueoflegends

import (
	"testing"

	"github.com/yuhanfang/riot/constants/champion"
)

var champions []champion.Champion

func addChampion(champ champion.Champion) []champion.Champion {
	champions = append(champions, champ)
	return champions
}

// TestProfileBuilder tests ProfileBuilder()
func TestProfileBuilder(t *testing.T) {/*
	var champ1 champion.Champion = 1
	var champ2 champion.Champion = 2
	var champ3 champion.Champion = 3

	champions = addChampion(champ1)
	champions = addChampion(champ2)
	champions = addChampion(champ3)

	profile := ProfileLOL{
		Scoring{
			Rank:    "SILVER I | 0 LP",
			Winrate: "54.01% W/L",
		},
		champions,
	}

	template := profile.ProfileBuilder()

	t.Run("Template is empty ?", func(t *testing.T) {
		if template == "" {
			t.Errorf("Template = %v, shouldn't be empty ", template)
		}
	})

	goodTemplate := "**SILVER I | 0 LP**\n"
	goodTemplate += "54.01% W/L\n\n"
	goodTemplate += "- **Champions : **Annie, Olaf, Galio"

	t.Run("Template has good values ?", func(t *testing.T) {
		if template != goodTemplate {
			t.Errorf("Template = %v, should be ", goodTemplate)
		}
	})
	*/
}
