package DockGo

func NewShard(token string, count int) []*Client {
	var total int
	shardList := make([]*Client, 0)
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
