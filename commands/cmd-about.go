package commands

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/intervinn/noorse"
)

var AboutCommand = &noorse.Command{
	Data: api.CreateCommandData{
		Name:        "about",
		Description: "about this bot",
	},
	Callback: func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		return EmbedResponse(discord.Embed{
			Title: "noorse",
			Description: `i exist to store your server points
			
			also i am [open sourced](https://github.com/intervinn/noorse) incase you want to update me

			i pay for a minecraft server and turns it provides 2 mysql instances
			so i have an almost infinite database as long as the server stays alive

			however i cant be arsed to buy some vps to host the actual bot
			theoretically i couldve written the bot in java and plug it in as a server plugin
			but yeah this bot is written in go
			`,
			Footer: &discord.EmbedFooter{
				Text: "that magic man made me",
			},
		})
	},
}

func init() {
	Commands = append(Commands, AboutCommand)
}
