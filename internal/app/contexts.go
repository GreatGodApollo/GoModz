package app

import (
	"github.com/GreatGodApollo/GoModz/pkg/api"
	"github.com/bwmarrin/discordgo"
	"io"
)

func (c CommandContext) Client() api.ClientWrapper {
	return c.client
}

func (c CommandContext) Message() *discordgo.Message {
	return c.message
}

func (c CommandContext) User() *discordgo.User {
	return c.user
}

func (c CommandContext) Channel() *discordgo.Channel {
	return c.channel
}

func (c CommandContext) Guild() *discordgo.Guild {
	return c.guild
}

func (c CommandContext) Member() *discordgo.Member {
	return c.member
}

func (c CommandContext) Reply(message string) (*discordgo.Message, error) {
	return c.client.GetSession().ChannelMessageSend(c.channel.ID, message)
}

func (c CommandContext) ReplyEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return c.client.GetSession().ChannelMessageSendEmbed(c.channel.ID, embed)
}

func (c CommandContext) ReplyEmbedBuilder(builder api.EmbedBuilder) (*discordgo.Message, error) {
	return c.client.GetSession().ChannelMessageSendEmbed(c.channel.ID, builder.Build())
}

func (c CommandContext) ReplyFile(filename string, file io.Reader) (*discordgo.Message, error) {
	return c.client.GetSession().ChannelFileSend(c.channel.ID, filename, file)
}

func (c CommandContext) ReplyTTS(message string) (*discordgo.Message, error) {
	return c.client.GetSession().ChannelMessageSendTTS(c.channel.ID, message)
}

func (m ModuleContext) Client() api.ClientWrapper {
	return m.client
}

func (m ModuleContext) Logger() api.Log {
	return m.log
}