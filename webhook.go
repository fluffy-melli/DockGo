package DockGo

import "github.com/bwmarrin/discordgo"

type SendWebhook struct {
	Text            string
	Files           []*discordgo.File
	Embeds          []*discordgo.MessageEmbed
	Buttons         []discordgo.Button
	SelectMenu      []discordgo.SelectMenu
	Attachments     []*discordgo.MessageAttachment
	AllowedMentions *discordgo.MessageAllowedMentions
	Ephemeral       bool
	ThreadName      string
	Username        string
	AvatarURL       string
	Wait            bool
	TTS             bool
}

type EditWebhook struct {
	Text            string
	Files           []*discordgo.File
	Embeds          []*discordgo.MessageEmbed
	Buttons         []discordgo.Button
	SelectMenu      []discordgo.SelectMenu
	Attachments     []*discordgo.MessageAttachment
	AllowedMentions discordgo.MessageAllowedMentions
}

type RespondWebhook discordgo.Message

func (rm *RespondWebhook) Method() *discordgo.Message {
	return (*discordgo.Message)(rm)
}
