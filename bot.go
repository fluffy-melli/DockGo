package DockGo

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var Logger = true

type Client discordgo.Session

func (bot *Client) Method() *discordgo.Session {
	return (*discordgo.Session)(bot)
}

func (bot *Client) Connect() {
	err := bot.Method().Open()
	if err != nil {
		Error(ERROR, "%v", err)
	}
}

func (bot *Client) Ready(function func(*Client, *Ready)) {
	go bot.Method().AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		cl := (*Client)(s)
		ry := (*Ready)(r)
		if Logger {
			ry.Logger(cl)
		}
		go function(cl, ry)
	})
}

func NewBot(token string) *Client {
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		Error(ERROR, "%v", err)
	}
	return (*Client)(bot)
}

func Wait() {
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-wait
}
