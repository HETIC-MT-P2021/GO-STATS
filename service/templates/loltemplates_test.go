package templates

import (
	"testing"

	"github.com/yuhanfang/riot/apiclient"
)

var champions []apiclient.ChampionMastery

func initThreeChampions() []apiclient.ChampionMastery {
	var champ1 apiclient.ChampionMastery = apiclient.ChampionMastery{
		ChampionID: 1,
	}
	var champ2 apiclient.ChampionMastery = apiclient.ChampionMastery{
		ChampionID: 2,
	}
	var champ3 apiclient.ChampionMastery = apiclient.ChampionMastery{
		ChampionID: 3,
	}

	champions = append(champions, champ1)
	champions = append(champions, champ2)
	champions = append(champions, champ3)

	return champions
}

func initEmptyChampions() []apiclient.ChampionMastery {
	return []apiclient.ChampionMastery{}
}

func TestProfileBuilderGoodTemplate(t *testing.T) {

	initThreeChampions()

	scoring := Scoring{
		Rank:    "SILVER I | 0 LP",
		Winrate: "54.01% W/L",
	}
	template := profileBuilderTemplate(scoring, champions)

	t.Run("Template is empty ?", func(t *testing.T) {
		if template == "" {
			t.Errorf("Template = %v, shouldn't be empty ", template)
		}
	})

	goodTemplate := "**SILVER I | 0 LP**\n"
	goodTemplate += "54.01% W/L\n\n"
	goodTemplate += "- **Champions : **\n\n"
	goodTemplate += " > :one: Annie - 0 pts\n"
	goodTemplate += " > :two: Olaf - 0 pts\n"
	goodTemplate += " > :three: Galio - 0 pts"

	t.Run("Template has good values ?", func(t *testing.T) {
		if template != goodTemplate {
			t.Errorf("Template = %v, should be %v", template, goodTemplate)
		}
	})
}

func TestProfileBuilderEmptyRank(t *testing.T) {

	initThreeChampions()

	scoring := Scoring{
		Rank:    "",
		Winrate: "",
	}
	template := profileBuilderTemplate(scoring, champions)

	t.Run("Template is empty ?", func(t *testing.T) {
		if template == "" {
			t.Errorf("Template = %v, shouldn't be empty ", template)
		}
	})

	goodTemplate := "Not ranked in Solo/Duo\n\n"
	goodTemplate += "- **Champions : **\n\n"
	goodTemplate += " > :one: Annie - 0 pts\n"
	goodTemplate += " > :two: Olaf - 0 pts\n"
	goodTemplate += " > :three: Galio - 0 pts"

	t.Run("Template has good values ?", func(t *testing.T) {
		if template != goodTemplate {
			t.Errorf("Template = \n%v, \nshould be \n%v", template, goodTemplate)
		}
	})
}

func TestProfileBuilderEmptyChampions(t *testing.T) {

	initEmptyChampions()

	scoring := Scoring{
		Rank:    "SILVER I | 0 LP",
		Winrate: "54.01% W/L",
	}
	template := profileBuilderTemplate(scoring, champions)

	t.Run("Template is empty ?", func(t *testing.T) {
		if template == "" {
			t.Errorf("Template = %v, shouldn't be empty ", template)
		}
	})

	goodTemplate := "**SILVER I | 0 LP**\n"
	goodTemplate += "54.01% W/L\n\n"
	goodTemplate += "- **Champions : **Not enough champions to display"

	t.Run("Template has good values ?", func(t *testing.T) {
		if template != goodTemplate {
			t.Errorf("Template = \n%v, should be \n%v", template, goodTemplate)
		}
	})
}

func profileBuilderTemplate(scoring Scoring, champions []apiclient.ChampionMastery) string {
	profile := ProfileLOL{
		scoring,
		champions,
	}

	return profile.ProfileBuilder()
}
