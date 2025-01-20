package DockGo

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type DockGo discordgo.Session

func (bot *DockGo) Method() *discordgo.Session {
	return (*discordgo.Session)(bot)
}

func (bot *DockGo) Connect() {
	err := bot.Method().Open()
	if err != nil {
		log.Fatalln(err)
	}
}

func (bot *DockGo) Ready(function func(*discordgo.Ready)) {
	bot.Method().AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		go function(r)
	})
}

func NewBot(token string) *DockGo {
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalln(err)
	}
	return (*DockGo)(bot)
}

func Wait() {
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-wait
}
