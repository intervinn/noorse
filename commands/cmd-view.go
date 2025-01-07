package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/intervinn/noorse"
	"github.com/intervinn/noorse/storage"
)

var ViewPointsCommand = &noorse.Command{
	Data: api.CreateCommandData{
		Name:        "view",
		Description: "view points",

		Options: discord.CommandOptions{
			&discord.UserOption{
				OptionName:  "user",
				Description: "user to view points of",
				Required:    false,
			},
		},
	},

	Callback: func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		guild, err := ParseGuild(data.Event.GuildID.String())
		if err != nil {
			return ErrorResponse("couldn't parse guild", err)
		}

		var user *discord.User = new(discord.User)
		option := data.Options.Find("user").String()
		if option == "" {
			user, err = ParseUser(data.Event.Sender().ID.String())
		} else {
			user, err = ParseUser(option)
		}

		if err != nil {
			return ErrorResponse("couldn't parse user", err)
		}

		a := new(storage.GuildAccount)
		err = storage.GetInstance().DB.Model(&storage.GuildAccount{
			UserID:  int64(user.ID),
			GuildID: int64(guild.ID),
		}).First(&a).Error
		if err != nil {
			return ErrorResponse("db query issue", err)
		}

		return EmbedResponse(discord.Embed{
			Title:       fmt.Sprintf("%s's points", user.DisplayName),
			Description: strconv.Itoa(int(a.Amount)),
		})
	},
}

func init() {
	Commands = append(Commands, ViewPointsCommand)
}
