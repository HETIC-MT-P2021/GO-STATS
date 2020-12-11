package structs

// DiscordEmbed Set up NewEmbed .SetThumbnail() || .SetTile() || .Description()
type DiscordEmbed struct {
	ProfileIconID int
	Title         DiscordEmbedTitle
	Description   string
}

// DiscordEmbedTitle Set up NewEmbed.SetTile()
type DiscordEmbedTitle struct {
	SummonerName  string
	SummonerLevel int64
}
