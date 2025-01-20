package main

import (
	"testing"

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

var Command = &DockGo.MessageCommands{
	Builder: &DockGo.MessageBuilder{
		Prefix:    "",
		StartWith: true,
	},
	Execute: func(client *DockGo.Client, mc *DockGo.MessageCreate) {
		ms := mc.SendMessage(client, &discordgo.MessageSend{})
		ms.EditMessage(client, &discordgo.MessageEdit{})
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
