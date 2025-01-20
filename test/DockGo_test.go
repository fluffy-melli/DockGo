package main

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/fluffy-melli/DockGo"
)

func TestBot(t *testing.T) {
	bot := DockGo.NewBot("~")
	bot.Ready(func(r *discordgo.Ready) {
		fmt.Println(r.User.Username + "가 실행이 되었습니다")
	})
	bot.Connect()
	DockGo.Wait()
}

func TestShard(t *testing.T) {
	shard := DockGo.NewShard("~", 2)
	for _, bot := range shard {
		bot.Ready(func(r *discordgo.Ready) {
			fmt.Println(r.User.Username + "가 실행이 되었습니다")
		})
		bot.Connect()
	}
	DockGo.Wait()
}

func TestMain(m *testing.M) {
	bot := DockGo.NewBot("~")
	bot.Ready(func(r *discordgo.Ready) {
		fmt.Println(r.User.Username + "가 실행이 되었습니다")
		bot.Register(&DockGo.MessageCommands{
			Builder: &DockGo.MessageBuilder{
				Name:      "",
				StartWith: true,
			},
			Execute: func(mc *DockGo.MessageCreate) {
				fmt.Println(mc.Method().Content)
			},
		})
		bot.Register(&DockGo.SlashCommands{
			Builder: &DockGo.SlashBuilder{
				Name:        "",
				Description: "",
			},
			Execute: func(i *DockGo.Interaction) {
				fmt.Println("")
			},
		})
	})
	bot.Connect()
	DockGo.Wait()
}
