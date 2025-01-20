package DockGo

import "github.com/bwmarrin/discordgo"

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
	Execute func(*Interaction)
}
