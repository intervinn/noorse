package commands

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/intervinn/noorse"
	"github.com/intervinn/noorse/storage"
)

var Commands []*noorse.Command = []*noorse.Command{}

func AddPoints(u *discord.User, g *discord.Guild) error {
	if !storage.GetInstance().UserExists(int64(u.ID)) {
		storage.GetInstance().DB.Create(&storage.User{
			ID: int64(u.ID),
			Accounts: []storage.GuildAccount{
				{
					ID:     int64(g.ID),
					Amount: 0,
				},
			},
		})
	}

	return nil
}

func EmbedResponse(embed discord.Embed) *api.InteractionResponseData {
	return &api.InteractionResponseData{
		Embeds: &[]discord.Embed{
			embed,
		},
	}
}

func ErrorResponse(err error) *api.InteractionResponseData {
	return EmbedResponse(ErrorEmbed(err))
}

func ErrorEmbed(err error) discord.Embed {
	return discord.Embed{
		Title:       "there was an issue",
		Description: err.Error(),
	}
}

func SuccessEmbed(user *discord.User, old int, new int) discord.Embed {
	return discord.Embed{
		Title: "af",
	}
}

func ParseUser(id string) (*discord.User, error) {
	s, err := discord.ParseSnowflake(id)
	if err != nil {
		return nil, err
	}
	u, err := noorse.GetInstance().State.User(discord.UserID(s))
	if err != nil {
		return nil, err
	}

	return u, nil
}
