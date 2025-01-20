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

func (mc *MessageCreate) SendMessage(client *Client, message *discordgo.MessageSend) *RespondMessage {
	return client.SendMessage(message, mc.Message.ChannelID)
}

func (mc *MessageCreate) ReplyMessage(client *Client, message *discordgo.MessageSend) *RespondMessage {
	message.Reference = mc.Method().Reference()
	return client.SendMessage(message, mc.Message.ChannelID)
}
