package DockGo

import "github.com/bwmarrin/discordgo"

type ButtonBuilder struct {
	CustomID string
}

type ButtonCreate struct {
	event  *discordgo.InteractionCreate
	client *Client
}

func (mc *ButtonCreate) Method() *discordgo.InteractionCreate {
	return mc.event
}

func (mc *ButtonCreate) Client() *Client {
	return mc.client
}

type ButtonCommands struct {
	Builder *ButtonBuilder
	Execute func(*ButtonCreate)
}

func (bc *ButtonCreate) Respond(message *discordgo.InteractionResponse) error {
	return bc.Client().Method().InteractionRespond(bc.Method().Interaction, message)
}
