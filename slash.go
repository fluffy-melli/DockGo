package DockGo

import (
	"github.com/bwmarrin/discordgo"
)

type SlashCreate struct {
	Options map[string]*discordgo.ApplicationCommandInteractionDataOption
	event   *discordgo.InteractionCreate
	client  *Client
}

func (it *SlashCreate) Method() *discordgo.InteractionCreate {
	return it.event
}

func (it *SlashCreate) Client() *Client {
	return it.client
}

type SlashBuilder discordgo.ApplicationCommand

func (sb *SlashBuilder) Method() *discordgo.ApplicationCommand {
	return (*discordgo.ApplicationCommand)(sb)
}

type SlashCommands struct {
	Builder *SlashBuilder
	Execute func(*SlashCreate)
}

func (mc *SlashCreate) Deferred() error {
	err := mc.client.Method().InteractionRespond(mc.Method().Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		Print(ERROR, "\033[41m\033[33m%v\033[0m", err)
		return err
	}
	return nil
}

func (mc *SlashCreate) Respond(respond *discordgo.InteractionResponse) error {
	err := mc.client.Method().InteractionRespond(mc.Method().Interaction, respond)
	if err != nil {
		Print(ERROR, "\033[41m\033[33m%v\033[0m", err)
		return err
	}
	return nil
}

func (mc *SlashCreate) SendMessage(message *discordgo.InteractionResponseData) error {
	err := mc.client.Method().InteractionRespond(mc.Method().Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: message,
	})
	if err != nil {
		Print(ERROR, "\033[41m\033[33m%v\033[0m", err)
		return err
	}
	return nil
}

func (mc *SlashCreate) EditMessage(message *discordgo.WebhookEdit) (*RespondMessage, error) {
	msg, err := mc.client.Method().InteractionResponseEdit(mc.Method().Interaction, message)
	if err != nil {
		Print(ERROR, "\033[41m\033[33m%v\033[0m", err)
		return nil, err
	}
	return (*RespondMessage)(msg), nil
}
