package commands

import (
	"context"
	"fmt"

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
		latency := noorse.GetInstance().State.Gateway().Latency()

		return &api.InteractionResponseData{
			Embeds: &[]discord.Embed{
				{
					Title:       fmt.Sprintf("%d ms", latency.Milliseconds()),
					Description: "this is the discord gateway latency\n\nthe database latency is probably even worse",
				},
			},
		}
	},
}

func init() {
	Commands = append(Commands, PingCommand)
}
