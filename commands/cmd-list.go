package commands

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/intervinn/noorse"
	"github.com/intervinn/noorse/storage"
)

func ListPointsClick(e *gateway.InteractionCreateEvent, data discord.ComponentInteraction) {
	msg := e.Message
	s := noorse.Instance().State

	if e.SenderID() != msg.Interaction.User.ID {
		return
	}

	page, err := strconv.Atoi(strings.TrimPrefix(string(data.ID()), "list-pg-"))
	if err != nil {
		fmt.Printf("failed strconv: %v\n", err)
		return
	}

	err = s.RespondInteraction(e.ID, e.Token, api.InteractionResponse{
		Type: api.UpdateMessage,
		Data: MakePage(page, e.GuildID),
	})

	if err != nil {
		fmt.Printf("failed to update msg: %v\n", err)
		return
	}
}

func MakePage(page int, guild discord.GuildID) *api.InteractionResponseData {
	s := noorse.Instance().State
	row := &discord.ActionRowComponent{
		&discord.ButtonComponent{
			Style:    discord.PrimaryButtonStyle(),
			CustomID: discord.ComponentID(fmt.Sprintf("list-pg-%v", page-1)),
			Label:    "back",
		},
		&discord.ButtonComponent{
			Style:    discord.PrimaryButtonStyle(),
			CustomID: discord.ComponentID(fmt.Sprintf("list-pg-%v", page+1)),
			Label:    "next",
		},
	}
	if page == 0 {
		row = &discord.ActionRowComponent{
			&discord.ButtonComponent{
				Style:    discord.PrimaryButtonStyle(),
				CustomID: discord.ComponentID(fmt.Sprintf("list-pg-%v", page+1)),
				Label:    "next",
			},
		}
	}

	var accs []storage.GuildAccount
	err := storage.Instance().Paginate(page).Where("guild_id = ? AND NOT amount = 0", guild).Find(&accs).Error
	if err != nil {
		return ErrorResponse("pagination failed", err)
	}

	desc := ""
	for i, a := range accs {
		m, err := s.Member(guild, discord.UserID(a.UserID))
		if err != nil {
			desc += fmt.Sprintf("`error: %v`\n", err)
			continue
		}

		desc += fmt.Sprintf("%d. %s (`%s`) - `%d`\n", i+1+(page*10), m.Nick, m.User.Username, a.Amount)
	}

	return &api.InteractionResponseData{
		Embeds: &[]discord.Embed{
			{
				Title:       fmt.Sprintf("Server - %v", guild),
				Description: desc,
			},
		},

		Components: discord.ComponentsPtr(row),
	}
}

var ListPointsCommand = &noorse.Command{
	Data: api.CreateCommandData{
		Name:        "list",
		Description: "list all current guild accounts",
	},

	Callback: func(ctx context.Context, data cmdroute.CommandData) *api.InteractionResponseData {
		return MakePage(0, data.Event.GuildID)
	},
}

func init() {
	Commands = append(Commands, ListPointsCommand)
}
