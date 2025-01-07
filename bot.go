package noorse

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
)

type Bot struct {
	State  *state.State
	Router *cmdroute.Router
}

func New() *Bot {
	token := os.Getenv("TOKEN")
	fmt.Println(token)

	r := cmdroute.NewRouter()
	s := state.New("Bot " + token)
	s.AddInteractionHandler(r)
	s.AddIntents(gateway.IntentGuildMessages)
	s.AddIntents(gateway.IntentMessageContent)

	return &Bot{
		State:  s,
		Router: r,
	}
}

func (b *Bot) Init(commands []*Command) {
	b.addFuncs(commands)
	b.overwriteCommands(commands)
}

func (b *Bot) addFuncs(commands []*Command) {
	for _, c := range commands {
		b.Router.AddFunc(c.Data.Name, (cmdroute.CommandHandlerFunc)(c.Callback))
	}
}

func (b *Bot) overwriteCommands(commands []*Command) {
	data := []api.CreateCommandData{}
	for _, c := range commands {
		data = append(data, c.Data)
	}

	if err := cmdroute.OverwriteCommands(b.State, data); err != nil {
		log.Fatalln("cannot update commands:", err)
	}
}

func (b *Bot) Serve() {
	if err := b.State.Connect(context.TODO()); err != nil {
		log.Println("cannot connect:", err)
	}
}
