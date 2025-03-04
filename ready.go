package DockGo

import (
	"github.com/bwmarrin/discordgo"
)

type Ready struct {
	event  *discordgo.Ready
	client *Client
}

func (mc *Ready) Method() *discordgo.Ready {
	return mc.event
}

func (mc *Ready) Client() *Client {
	return mc.client
}

func (mc *Ready) Logger(client *Client) {
	shard := client.Method().ShardCount
	if shard != 0 {
		Print(INFO, "\033[0m\033[30m\033[42mReady!\033[0m [\033[34mID: %d\033[0m / \033[33mTotal: %d\033[0m]\n", client.Method().ShardID, shard)
	} else {
		Print(INFO, "\033[30m\033[42mReady!\033[0m \033[41m\033[33m(No shards)\033[0m")
	}
}
