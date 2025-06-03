package commands

import (
	"github.com/purrquinox/zevola/types"
	"github.com/purrquinox/zevola/core"
)

func EnableDev(evt types.Event, stdout *bool) {
	*stdout = true
	core.Respond(evt, "Enabled developer mode.", nil, nil)
}

func DisableDev(evt types.Event, stdout *bool) {
	*stdout = false
	core.Respond(evt, "Disabled developer mode.", nil, nil)
}
