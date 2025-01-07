package commands

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/intervinn/noorse"
)

var PingCommand = &noorse.Command{
	Data: api.CreateCommandData{
		Name:        "ping",
		Description: "ping",
	},
	Callback: func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		return &api.InteractionResponseData{
			Embeds: &[]discord.Embed{
				{
					Title: "hi",
				},
			},
		}
	},
}

func init() {
	Commands = append(Commands, PingCommand)
}
