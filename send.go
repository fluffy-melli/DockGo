package DockGo

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type RespondMessage discordgo.Message

func (rm *RespondMessage) Method() *discordgo.Message {
	return (*discordgo.Message)(rm)
}

func (bot *Client) SendMessage(message *discordgo.MessageSend, channelID string) *RespondMessage {
	msg, err := bot.Method().ChannelMessageSendComplex(channelID, message)
	if err != nil {
		log.Println("error sending complex message,", err)
	}
	return (*RespondMessage)(msg)
}

func (bot *Client) EditMessage(message *discordgo.MessageEdit, channelID, messageID string) *RespondMessage {
	msg, err := bot.Method().ChannelMessageEditComplex(message)
	if err != nil {
		log.Println("error editing complex message,", err)
	}
	return (*RespondMessage)(msg)
}

func (bot *Client) DeleteMessage(channelID, messageID string) {
	err := bot.Method().ChannelMessageDelete(channelID, messageID)
	if err != nil {
		log.Println("error deleteing complex message,", err)
	}
}

func (rm *RespondMessage) Followup(client *Client, message *discordgo.MessageSend) *RespondMessage {
	message.Reference = client.Reference(rm.Method().GuildID, rm.Method().ChannelID, rm.Method().ID)
	return client.SendMessage(message, rm.Method().ChannelID)
}

func (rm *RespondMessage) EditMessage(client *Client, message *discordgo.MessageEdit) *RespondMessage {
	return client.EditMessage(message, rm.Method().ChannelID, rm.Method().ID)
}

func (rm *RespondMessage) DeleteMessage(client *Client) {
	go client.DeleteMessage(rm.Method().ChannelID, rm.Method().ID)
}

func (bot *Client) Reference(guildID string, channelID string, messageID string) *discordgo.MessageReference {
	return &discordgo.MessageReference{
		MessageID: messageID,
		ChannelID: channelID,
		GuildID:   guildID,
	}
}
