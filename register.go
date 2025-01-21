package DockGo

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (bot *Client) Register(command interface{}) {
	switch command := command.(type) {
	case MessageCommands:
	case *MessageCommands:
		go bot.Method().AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
			if m.Content == command.Builder.Prefix {
				go command.Execute((*Client)(s), (*MessageCreate)(m))
			} else if command.Builder.StartWith && strings.HasPrefix(m.Content, command.Builder.Prefix) {
				go command.Execute((*Client)(s), (*MessageCreate)(m))
			}
		})
	case SlashCommands:
	case *SlashCommands:
		_, err := bot.Method().ApplicationCommandCreate(bot.Method().State.User.ID, "", command.Builder.Method())
		if err != nil {
			Error(ERROR, "\033[41m\033[33m%v\033[0m", err)
		}
		go bot.Method().AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type == discordgo.InteractionMessageComponent {
				return
			}
			if i.ApplicationCommandData().Name == command.Builder.Name {
				go command.Execute((*Client)(s), (*InteractionCreate)(i))
			}
		})
	default:
		Error(ERROR, "\033[41m\033[33munknown type : %v\033[0m", command)
	}
}
