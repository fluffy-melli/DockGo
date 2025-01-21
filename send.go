package DockGo

import (
	"github.com/bwmarrin/discordgo"
)

type RespondMessage discordgo.Message

func (rm *RespondMessage) Method() *discordgo.Message {
	return (*discordgo.Message)(rm)
}

func (bot *Client) SendMessage(message *discordgo.MessageSend, channelID string) (*RespondMessage, error) {
	msg, err := bot.Method().ChannelMessageSendComplex(channelID, message)
	if err != nil {
		Print(ERROR, "\033[41m\033[33m%v\033[0m", err)
		return nil, err
	}
	return (*RespondMessage)(msg), err
}

func (bot *Client) EditMessage(message *discordgo.MessageEdit, channelID, messageID string) (*RespondMessage, error) {
	msg, err := bot.Method().ChannelMessageEditComplex(message)
	if err != nil {
		Print(ERROR, "\033[41m\033[33m%v\033[0m", err)
		return nil, err
	}
	return (*RespondMessage)(msg), nil
}

func (bot *Client) DeleteMessage(channelID, messageID string) error {
	err := bot.Method().ChannelMessageDelete(channelID, messageID)
	if err != nil {
		Print(ERROR, "\033[41m\033[33m%v\033[0m", err)
		return err
	}
	return nil
}

func (rm *RespondMessage) Followup(client *Client, message *discordgo.MessageSend) (*RespondMessage, error) {
	message.Reference = client.Reference(rm.Method().GuildID, rm.Method().ChannelID, rm.Method().ID)
	msg, err := client.SendMessage(message, rm.Method().ChannelID)
	if err != nil {
		Print(ERROR, "\033[41m\033[33m%v\033[0m", err)
		return nil, err
	}
	return (*RespondMessage)(msg), nil
}

func (rm *RespondMessage) EditMessage(client *Client, message *discordgo.MessageEdit) (*RespondMessage, error) {
	msg, err := client.EditMessage(message, rm.Method().ChannelID, rm.Method().ID)
	if err != nil {
		Print(ERROR, "\033[41m\033[33m%v\033[0m", err)
		return nil, err
	}
	return (*RespondMessage)(msg), nil
}

func (rm *RespondMessage) DeleteMessage(client *Client) error {
	return client.DeleteMessage(rm.Method().ChannelID, rm.Method().ID)
}

func (bot *Client) Reference(guildID string, channelID string, messageID string) *discordgo.MessageReference {
	return &discordgo.MessageReference{
		MessageID: messageID,
		ChannelID: channelID,
		GuildID:   guildID,
	}
}
