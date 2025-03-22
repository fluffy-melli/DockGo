package DockGo

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (bot *Client) Register(command interface{}) {
	switch command := command.(type) {
	case *MessageCommands:
		go bot.Method().AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
			if m.Content == command.Builder.Prefix {
				go command.Execute(&MessageCreate{
					event:  m,
					client: (*Client)(s),
				})
			} else if command.Builder.StartWith && strings.HasPrefix(m.Content, command.Builder.Prefix) {
				go command.Execute(&MessageCreate{
					event:  m,
					client: (*Client)(s),
				})
			}
		})
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
				var options = make(map[string]*discordgo.ApplicationCommandInteractionDataOption)
				for _, option := range i.ApplicationCommandData().Options {
					options[option.Name] = option
				}
				go command.Execute(&SlashCreate{
					event:   i,
					client:  (*Client)(s),
					Options: options,
				})
			}
		})
	case *ButtonCommands:
		go bot.Method().AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type == discordgo.InteractionApplicationCommand {
				return
			}
			if i.Type == discordgo.InteractionMessageComponent && i.MessageComponentData().CustomID == command.Builder.CustomID {
				go command.Execute(&ButtonCreate{
					event:  i,
					client: (*Client)(s),
				})
			}
		})
	default:
		Error(ERROR, "\033[41m\033[33munknown type : %v\033[0m", command)
	}
}
