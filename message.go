package DockGo

import (
	"github.com/bwmarrin/discordgo"
)

type MessageBuilder struct {
	Prefix    string
	StartWith bool
}

type MessageCreate discordgo.MessageCreate

func (mc *MessageCreate) Method() *discordgo.MessageCreate {
	return (*discordgo.MessageCreate)(mc)
}

type MessageCommands struct {
	Builder *MessageBuilder
	Execute func(*Client, *MessageCreate)
}
