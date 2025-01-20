package DockGo

import (
	"log"

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

func (mc *InteractionCreate) SendMessage(client *Client, message *discordgo.WebhookEdit) *RespondMessage {
	msg, err := client.Method().InteractionResponseEdit(mc.Method().Interaction, message)
	if err != nil {
		log.Println("error sending complex message,", err)
	}
	return (*RespondMessage)(msg)
}
