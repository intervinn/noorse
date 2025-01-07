package commands

import (
	"context"
	"strings"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/intervinn/noorse"
)

var AddPointsCommand = &noorse.Command{
	Data: api.CreateCommandData{
		Name:        "add",
		Description: "add points",
		Options: discord.CommandOptions{
			&discord.IntegerOption{
				Required:    true,
				OptionName:  "amount",
				Description: "the amount of points to add",
			},
			&discord.StringOption{
				Required:    true,
				OptionName:  "reason",
				Description: "specify a reason for the points assignment",
			},
			&discord.UserOption{
				Required:    false,
				OptionName:  "user",
				Description: "user to add points to, to select multiple use `userids`",
			},
			&discord.StringOption{
				Required:    false,
				OptionName:  "userids",
				Description: "write user ids splitted by space to assign points to all of them",
			},
		},
	},
	Callback: func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		guild, err := ParseGuild(data.Event.GuildID.String())
		if err != nil {
			return ErrorResponse("couldn't parse guild", err)
		}

		user := data.Data.Options.Find("user").String()
		userids := data.Data.Options.Find("userids").String()

		if userids == "" && user == "" {
			return EmbedResponse(discord.Embed{
				Title: "specify either a user or userids duh",
			})
		}

		amount, err := data.Data.Options.Find("amount").IntValue()
		if err != nil {
			return ErrorResponse("couldn't parse amount", err)
		}

		ids := strings.Split(userids, " ")
		if user != "" {
			ids = append(ids, user)
		}

		embeds := []discord.Embed{}
		// convert to users
		for _, i := range ids {
			if i == "" {
				continue
			}

			u, err := ParseUser(i)
			if err != nil {
				embeds = append(embeds, ErrorEmbed("failed to parse user", err))
				continue
			}

			prev, new, err := AddPoints(u, guild, amount)
			if err != nil {
				embeds = append(embeds, ErrorEmbed("failed to add points", err))
				continue
			}

			embeds = append(embeds, SuccessEmbed(u, prev, new))
		}

		return &api.InteractionResponseData{
			Embeds: &embeds,
		}
	},
}

func init() {
	Commands = append(Commands, AddPointsCommand)
}
