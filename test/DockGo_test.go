package main

import (
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/fluffy-melli/DockGo"
)

func TestBot(t *testing.T) {
	bot := DockGo.NewBot("~")
	bot.Ready(func(_ *DockGo.Client, r *DockGo.Ready) {})
	bot.Connect()
	DockGo.Wait()
}

func TestShard(t *testing.T) {
	shard := DockGo.NewShard("~", 2)
	for _, bot := range shard {
		bot.Ready(func(_ *DockGo.Client, r *DockGo.Ready) {})
		bot.Connect()
	}
	DockGo.Wait()
}

var Command = &DockGo.SlashCommands{
	Builder: &DockGo.SlashBuilder{
		Name:        "야옹",
		Description: "애오옹",
	},
	Execute: func(c *DockGo.Client, ic *DockGo.InteractionCreate) {
		ic.SendMessage(c, &discordgo.InteractionResponseData{
			Content: "??",
		})
		time.Sleep(3 * time.Second)
		var text = "?!!"
		ic.EditMessage(c, &discordgo.WebhookEdit{
			Content: &text,
		})
	},
}

func TestMain(m *testing.M) {
	shard := DockGo.NewShard("~", 2)
	for _, bot := range shard {
		bot.Ready(func(_ *DockGo.Client, r *DockGo.Ready) {
			bot.Register(Command)
		})
		bot.Connect()
	}
	DockGo.Wait()
}
