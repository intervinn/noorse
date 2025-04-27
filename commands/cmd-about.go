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
			
			noorse is [open sourced](https://github.com/intervinn/noorse), feel free to contribute
			
			the database is provided from a minecraft server hosting

			however i cba to buy an actual hosting to deploy the bot, so it runs on some free thing

			to add and remove points a user must have a role named exactly 'Bot Manager', that's all you need
			`,
			Footer: &discord.EmbedFooter{
				Text: "there lived a certain man in russia long ago, he was cool and shit and he made noorse",
			},
		})
	},
}

func init() {

	Commands = append(Commands, AboutCommand)
}
