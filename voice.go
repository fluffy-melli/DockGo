package DockGo

import "github.com/bwmarrin/discordgo"

func (ic *SlashCreate) GetVoiceState() *discordgo.VoiceState {
	voiceState, err := ic.Client().State.VoiceState(ic.Method().GuildID, ic.Method().Member.User.ID)
	if err != nil {
		return nil
	}
	return voiceState
}

func (ic *SlashCreate) GetVoiceChannel() *discordgo.Channel {
	voiceState := ic.GetVoiceState()
	if voiceState == nil {
		return nil
	}
	channel, err := ic.Client().State.Channel(voiceState.ChannelID)
	if err != nil {
		return nil
	}
	return channel
}

func (ic *SlashCreate) GetVoiceConnection() *discordgo.VoiceConnection {
	return ic.Client().GetVoiceConnection(ic.Method().GuildID)
}

func (ic *SlashCreate) JoinVoiceChannel(channelID string) (*discordgo.VoiceConnection, error) {
	return ic.Client().JoinVoiceChannel(ic.Method().GuildID, channelID)
}

func (ic *SlashCreate) LeaveVoiceChannel() bool {
	return ic.Client().LeaveVoiceChannel(ic.Method().GuildID)
}

func (mc *MessageCreate) GetVoiceState() *discordgo.VoiceState {
	voiceState, err := mc.Client().State.VoiceState(mc.Method().GuildID, mc.Method().Author.ID)
	if err != nil {
		return nil
	}
	return voiceState
}

func (mc *MessageCreate) GetVoiceChannel() *discordgo.Channel {
	voiceState := mc.GetVoiceState()
	if voiceState == nil {
		return nil
	}
	channel, err := mc.Client().State.Channel(voiceState.ChannelID)
	if err != nil {
		return nil
	}
	return channel
}

func (mc *MessageCreate) GetVoiceConnection() *discordgo.VoiceConnection {
	return mc.Client().GetVoiceConnection(mc.Method().GuildID)
}

func (mc *MessageCreate) JoinVoiceChannel(channelID string) (*discordgo.VoiceConnection, error) {
	return mc.Client().JoinVoiceChannel(mc.Method().GuildID, channelID)
}

func (mc *MessageCreate) LeaveVoiceChannel() bool {
	return mc.Client().LeaveVoiceChannel(mc.Method().GuildID)
}

func (bot *Client) JoinVoiceChannel(guildID, channelID string) (*discordgo.VoiceConnection, error) {
	vs, err := bot.Method().ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		Print(ERROR, "%v", err)
		return nil, err
	}
	return vs, nil
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
