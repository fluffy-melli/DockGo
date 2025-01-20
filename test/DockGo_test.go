package main

import (
	"fmt"
	"testing"

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

func TestMain(m *testing.M) {
	shard := DockGo.NewShard("~", 2)
	for _, bot := range shard {
		bot.Ready(func(_ *DockGo.Client, r *DockGo.Ready) {
			bot.Register(&DockGo.MessageCommands{
				Builder: &DockGo.MessageBuilder{
					Prefix:    "",
					StartWith: true,
				},
				Execute: func(_ *DockGo.Client, mc *DockGo.MessageCreate) {
					fmt.Println(mc.Method().Content)
				},
			})
		})
		bot.Connect()
	}
	DockGo.Wait()
}
