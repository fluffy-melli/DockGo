package DockGo

import (
	"github.com/bwmarrin/discordgo"
)

type MessageBuilder struct {
	Prefix    string
	StartWith bool
}

type MessageCreate struct {
	event  *discordgo.MessageCreate
	client *Client
}

func (mc *MessageCreate) Method() *discordgo.MessageCreate {
	return mc.event
}

func (mc *MessageCreate) Client() *Client {
	return mc.client
}

type MessageCommands struct {
	Builder *MessageBuilder
	Execute func(*MessageCreate)
}

func (mc *MessageCreate) SendMessage(message *discordgo.MessageSend) (*RespondMessage, error) {
	return mc.client.SendMessage(message, mc.event.Message.ChannelID)
}

func (mc *MessageCreate) ReplyMessage(message *discordgo.MessageSend) (*RespondMessage, error) {
	message.Reference = mc.Method().Reference()
	return mc.client.SendMessage(message, mc.event.Message.ChannelID)
}
