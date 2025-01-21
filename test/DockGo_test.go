package main

import (
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/fluffy-melli/DockGo"
)

func TestBot(t *testing.T) {
	bot := DockGo.NewBot("~")
	bot.Ready(func(r *DockGo.Ready) {})
	bot.Connect()
	DockGo.Wait()
}

func TestShard(t *testing.T) {
	shard := DockGo.NewShard("~", 2)
	for _, bot := range shard {
		bot.Ready(func(r *DockGo.Ready) {})
		bot.Connect()
	}
	DockGo.Wait()
}

var Command = &DockGo.SlashCommands{
	Builder: &DockGo.SlashBuilder{
		Name:        "야옹",
		Description: "애오옹",
	},
	Execute: func(ic *DockGo.SlashCreate) {
		ic.SendMessage(&discordgo.InteractionResponseData{
			Content: "??",
		})
		time.Sleep(3 * time.Second)
		var text = "?!!"
		ic.EditMessage(&discordgo.WebhookEdit{
			Content: &text,
		})
	},
}

func TestMain(m *testing.M) {
	shard := DockGo.NewShard("~", 2)
	for _, bot := range shard {
		bot.Ready(func(r *DockGo.Ready) {
			bot.Register(Command)
		})
		bot.Connect()
	}
	DockGo.Wait()
}
