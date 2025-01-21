package DockGo

import "github.com/bwmarrin/discordgo"

func (bot *Client) JoinVoiceChannel(guildID, channelID string) error {
	_, err := bot.Method().ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		Print(ERROR, "%v", err)
		return err
	}
	return nil
}

func (bot *Client) LeaveVoiceChannel(guildID string) bool {
	for _, vs := range bot.Method().VoiceConnections {
		if vs.GuildID == guildID {
			vs.Disconnect()
			return true
		}
	}
	return false
}

func (bot *Client) GetVoiceConnection(guildID string) *discordgo.VoiceConnection {
	for _, vs := range bot.Method().VoiceConnections {
		if vs.GuildID == guildID {
			return vs
		}
	}
	return nil
}
