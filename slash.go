package DockGo

import (
	"github.com/bwmarrin/discordgo"
)

type InteractionCreate discordgo.InteractionCreate

func (it *InteractionCreate) Method() *discordgo.InteractionCreate {
	return (*discordgo.InteractionCreate)(it)
}

type SlashBuilder discordgo.ApplicationCommand

func (sb *SlashBuilder) Method() *discordgo.ApplicationCommand {
	return (*discordgo.ApplicationCommand)(sb)
}

type SlashCommands struct {
	Builder *SlashBuilder
	Execute func(*Client, *InteractionCreate)
}

func (mc *InteractionCreate) Deferred(client *Client) {
	go client.Method().InteractionRespond(mc.Method().Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
}

func (mc *InteractionCreate) Respond(client *Client, respond *discordgo.InteractionResponse) {
	go client.Method().InteractionRespond(mc.Method().Interaction, respond)
}

func (mc *InteractionCreate) SendMessage(client *Client, message *discordgo.InteractionResponseData) {
	go client.Method().InteractionRespond(mc.Method().Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: message,
	})
}

func (mc *InteractionCreate) EditMessage(client *Client, message *discordgo.WebhookEdit) *RespondMessage {
	msg, err := client.Method().InteractionResponseEdit(mc.Method().Interaction, message)
	if err != nil {
		Print(ERROR, "\033[41m\033[33m%v\033[0m", err)
	}
	return (*RespondMessage)(msg)
}
