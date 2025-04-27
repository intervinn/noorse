package commands

import (
	"fmt"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/intervinn/noorse"
	"github.com/intervinn/noorse/storage"
)

var Commands []*noorse.Command = []*noorse.Command{}

func Init() {
	noorse.Instance().State.AddHandler(func(e *gateway.InteractionCreateEvent) {
		switch data := e.Data.(type) {
		case discord.ComponentInteraction:
			ListPointsClick(e, data)
			return
		}
	})
}

func IsDev(u *discord.User) bool {
	return u.ID == 347365756301737994
}

func IsManager(m *discord.Member, guild *discord.Guild) bool {
	state := noorse.Instance().State

	for _, rid := range m.RoleIDs {
		role, err := state.Role(guild.ID, rid)
		if err != nil {
			continue
		}

		if role.Name == "Bot Manager" {
			return true
		}
	}
	return false
}

func AddPoints(u *discord.User, g *discord.Guild, amount int64) (int64, int64, error) {
	if !storage.Instance().UserExists(int64(g.ID), int64(u.ID)) {
		fmt.Println("USER DOESNT EXIST, CREATE")
		storage.Instance().DB.Create(&storage.GuildAccount{
			UserID:  int64(u.ID),
			GuildID: int64(g.ID),
			Amount:  0,
		})
	}

	record := new(storage.GuildAccount)
	err := storage.Instance().DB.Where("guild_id = ? AND user_id = ?", g.ID, u.ID).First(record).Error
	if err != nil {
		return 0, 0, err
	}

	prev := record.Amount
	new := record.Amount + amount

	record.Amount = new
	if err := storage.Instance().DB.Save(record).Error; err != nil {
		return 0, 0, err
	}

	return prev, new, nil
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

func ParseUser(id string) (*discord.User, error) {
	s, err := discord.ParseSnowflake(id)
	if err != nil {
		return nil, err
	}
	u, err := noorse.Instance().State.User(discord.UserID(s))
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

	g, err := noorse.Instance().State.Guild(discord.GuildID(s))
	if err != nil {
		return nil, err
	}
	return g, err
}
