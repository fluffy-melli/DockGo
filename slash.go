package DockGo

import (
	"github.com/bwmarrin/discordgo"
)

type SlashCreate struct {
	Options map[string]any
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

func (mc *SlashCreate) Deferred() {
	go mc.client.Method().InteractionRespond(mc.Method().Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
}

func (mc *SlashCreate) Respond(respond *discordgo.InteractionResponse) {
	go mc.client.Method().InteractionRespond(mc.Method().Interaction, respond)
}

func (mc *SlashCreate) SendMessage(message *discordgo.InteractionResponseData) {
	go mc.client.Method().InteractionRespond(mc.Method().Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: message,
	})
}

func (mc *SlashCreate) EditMessage(message *discordgo.WebhookEdit) *RespondMessage {
	msg, err := mc.client.Method().InteractionResponseEdit(mc.Method().Interaction, message)
	if err != nil {
		Print(ERROR, "\033[41m\033[33m%v\033[0m", err)
	}
	return (*RespondMessage)(msg)
}
