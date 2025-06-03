package commands

import (
	"fmt"

	"github.com/purrquinox/zevola/core"
	"github.com/purrquinox/zevola/types"
)

func TestEmbed(evt types.Event, _ *bool) {
	_, err := core.Respond(evt, "", &types.Embed{
		Title:       "Test Embed",
		Description: "This is a test embed.",
		URL:         core.Ptr("https://purrquinox.com/"),
		IconURL:     core.Ptr("https://purrquinox.com/logo.png"),
		Fields: core.Ptr([]types.EmbedField{
			{
				Name:  "Test Field",
				Value: "This is a test field.",
			},
		}),
		Footer: &types.EmbedFooter{
			Text:     "This is a test footer.",
			PhotoURL: "https://purrquinox.com/logo.png",
		},
		Color: 0x00FF00,
	}, nil)
	if err != nil {
		fmt.Printf("Error sending embed: %v\n", err)
	}
}
