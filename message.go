package DockGo

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type MessageBuilder struct {
	Name      string
	StartWith bool
}

type MessageCreate discordgo.MessageCreate

func (mc *MessageCreate) Method() *discordgo.MessageCreate {
	return (*discordgo.MessageCreate)(mc)
}

type Message struct {
	Builder MessageBuilder
	Execute func(*MessageCreate)
}

func (bot *DockGo) Register(command interface{}) {
	switch command := command.(type) {
	case Message:
	case *Message:
		bot.Method().AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
			if m.Content == command.Builder.Name {
				command.Execute((*MessageCreate)(m))
			} else if command.Builder.StartWith && strings.HasPrefix(m.Content, command.Builder.Name) {
				command.Execute((*MessageCreate)(m))
			}
		})
	default:
		log.Fatalln("unknow type")
	}
}
