package commands

import (
	"fmt"

	"github.com/purrquinox/zevola/types"
	"github.com/purrquinox/zevola/core"
	"github.com/bwmarrin/discordgo"
)

var commandMap = map[string]func(types.Event, *bool){
	"ping":        Ping,
	"test":        Test,
	"test_embed":  TestEmbed,
	"enable_dev":  EnableDev,
	"disable_dev": DisableDev,
}

func Handle(evt types.Event, stdout *bool, name string) {
	if handler, ok := commandMap[name]; ok {
		handler(evt, stdout)
	} else if *stdout {
		core.Respond(evt, "", &types.Embed{
			Title:       "types.Event Received",
			Description: fmt.Sprintf("%+v", evt),
			Color:       0x00FF00,
		}, nil)
	}
}

func Definitions() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Check latency",
		},
		{
			Name:        "test",
			Description: "Test types.event reply",
		},
		{
			Name:        "test_embed",
			Description: "Send a test embed",
		},
		{
			Name:        "enable_dev",
			Description: "Enable developer mode",
		},
		{
			Name:        "disable_dev",
			Description: "Disable developer mode",
		},
	}
}
