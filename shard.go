package DockGo

type Shard struct {
	Total int
	ID    int
	Bot   *DockGo
}

func NewShard(token string, count int) []*DockGo {
	var total int
	shardList := make([]*DockGo, 0)
	if count == 0 {
		bot := NewBot(token)
		total = (len(bot.Method().State.Guilds) + 999) / 1000
	} else {
		total = count - 1
	}
	for i := 0; i <= total; i++ {
		bot := NewBot(token)
		bot.Method().ShardCount = total
		bot.Method().ShardID = i
		shardList = append(shardList, bot)
	}
	return shardList
}
