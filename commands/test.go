package commands

import (
	"github.com/purrquinox/zevola/core"
	"github.com/purrquinox/zevola/types"
)

func Test(evt types.Event, _ *bool) {
	core.Respond(evt, "Received test event!", nil, nil)
}
