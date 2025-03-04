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
		Error(ERROR, "\033[41m\033[33m%v\033[0m", err)
	}
}

func (bot *Client) Ready(function func(*Ready)) {
	go bot.Method().AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		respond := &Ready{
			event:  r,
			client: (*Client)(s),
		}
		if Logger {
			respond.Logger(respond.client)
		}
		go function(respond)
	})
}

func NewBot(token string) *Client {
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		Error(ERROR, "\033[41m\033[33m%v\033[0m", err)
	}
	return (*Client)(bot)
}

func Wait() {
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-wait
}
