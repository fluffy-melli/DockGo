package DockGo

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Interaction discordgo.InteractionCreate

func (it *Interaction) Method() *discordgo.InteractionCreate {
	return (*discordgo.InteractionCreate)(it)
}

type SlashBuilder discordgo.ApplicationCommand

func (sb *SlashBuilder) Method() *discordgo.ApplicationCommand {
	return (*discordgo.ApplicationCommand)(sb)
}

type SlashCommands struct {
	Builder *SlashBuilder
	Execute func(*Client, *Interaction)
}

func (mc *Interaction) SendMessage(client *Client, message *discordgo.WebhookEdit) *RespondMessage {
	msg, err := client.Method().InteractionResponseEdit(mc.Method().Interaction, message)
	if err != nil {
		log.Println("error sending complex message,", err)
	}
	return (*RespondMessage)(msg)
}
