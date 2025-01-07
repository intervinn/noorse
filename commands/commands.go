package commands

import (
	"fmt"
	"strconv"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/intervinn/noorse"
	"github.com/intervinn/noorse/storage"
)

var Commands []*noorse.Command = []*noorse.Command{}

func AddPoints(u *discord.User, g *discord.Guild, amount int64) (int64, int64, error) {
	if !storage.GetInstance().UserExists(int64(g.ID), int64(u.ID)) {
		fmt.Println("NO EXSIT")
		storage.GetInstance().DB.Create(&storage.GuildAccount{
			UserID:  int64(u.ID),
			GuildID: int64(g.ID),
			Amount:  0,
		})
	}

	record := new(storage.GuildAccount)
	err := storage.GetInstance().DB.Where("guild_id = ? AND user_id = ?", g.ID, u.ID).First(record).Error
	if err != nil {
		return 0, 0, err
	}

	prev := record.Amount
	new := record.Amount + amount

	record.Amount = new
	if err := storage.GetInstance().DB.Save(record).Error; err != nil {
		return 0, 0, err
	}

	return prev, new, nil
}

func GetPoints(u *discord.User, g *discord.User) {

}

func EmbedResponse(embed discord.Embed) *api.InteractionResponseData {
	return &api.InteractionResponseData{
		Embeds: &[]discord.Embed{
			embed,
		},
	}
}

func ErrorResponse(message string, err error) *api.InteractionResponseData {
	return EmbedResponse(ErrorEmbed(message, err))
}

func ErrorEmbed(message string, err error) discord.Embed {
	return discord.Embed{
		Title:       "there was an issue",
		Description: fmt.Sprintf("%s: %s", message, err.Error()),
	}
}

func SuccessEmbed(user *discord.User, old int64, new int64) discord.Embed {
	return discord.Embed{
		Title: "af",
		Fields: []discord.EmbedField{
			{
				Name:   "old",
				Value:  strconv.Itoa(int(old)),
				Inline: false,
			},
			{
				Name:   "new",
				Value:  strconv.Itoa(int(new)),
				Inline: false,
			},
		},
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

func ParseGuild(id string) (*discord.Guild, error) {
	s, err := discord.ParseSnowflake(id)
	if err != nil {
		return nil, err
	}

	g, err := noorse.GetInstance().State.Guild(discord.GuildID(s))
	if err != nil {
		return nil, err
	}
	return g, err
}
