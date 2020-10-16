# GO-STATS

## What is GO-STATS ?
It's a Discord bot which allows you to get statistics of a video game about people in a discord chanel. In a first time, you will be able to get statistics about player(s) playing League of Legend exclusively.

## Features

- Get statistics about a specific player 


`-gs stats <game> <username>`
```http
GET /stats/{player_id} 
```

- Get statistics about all players


`-gs stats <game>`
```http
GET /stats 
```

## Libraries
- https://gowalker.org/github.com/bwmarrin/discordgo
- https://godoc.org/github.com/atuleu/go-lol

## Architecture of project
https://github.com/HETIC-MT-P2021/GO-STATS/tree/main/docs

## License  
[MIT](https://github.com/HETIC-MT-P2021/GO-STATS/blob/main/LICENSE)
