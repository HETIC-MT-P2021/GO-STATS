package service

import (
	"fmt"

	"github.com/yuhanfang/riot/constants/champion"
)

type ProfileLOL struct {
	Rank          string
	Winrate       string
	Champions     []champion.Champion
}

func (profile ProfileLOL) ProfileBuilder() string {
	template := fmt.Sprintf("**%s**\n%s\n", profile.Rank, profile.Winrate)
	template += "\n- **Champions : **"

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
