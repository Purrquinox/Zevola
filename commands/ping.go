package commands

import (
	"fmt"
	"time"

	"github.com/purrquinox/zevola/core"
	"github.com/purrquinox/zevola/types"

	"github.com/bwmarrin/discordgo"
	"github.com/sentinelb51/revoltgo"
)

func Ping(evt types.Event, _ *bool) {
	start := time.Now()
	msgRaw, err := core.Respond(evt, "Pinging...", nil, nil)
	if err != nil {
		fmt.Printf("Error sending ping: %v\n", err)
		return
	}
	latency := time.Since(start).Milliseconds()
	pong := fmt.Sprintf("Pong! %dms", latency)
	switch msg := msgRaw.(type) {
	case *discordgo.Message:
		core.Respond(evt, pong, nil, &msg.ID)
	case *revoltgo.Message:
		core.Respond(evt, pong, nil, &msg.ID)
	}
}
