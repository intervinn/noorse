package noorse

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
)

// prevent cyclic import
type Command struct {
	Data     api.CreateCommandData
	Callback cmdroute.CommandHandlerFunc
}
