package app

import (
	"errors"
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

// PurgeMessages purges 'x' number of messages from the Channel a CommandContext was initiated for.
func (c CommandContext) PurgeMessages(num int) error {
	if num >= 1 && num <= 100 {
		msgs, err := c.Client().GetSession().ChannelMessages(c.Channel().ID, num, "", "", "")
		if err != nil {
			return err
		}
		var ids []string
		for _, msg := range msgs {
			ids = append(ids, msg.ID)
		}
		return c.Client().GetSession().ChannelMessagesBulkDelete(c.Channel().ID, ids)
	} else if num > 1 && num > 100 {
		return errors.New("too many messages")
	} else if num == 0 {
		return errors.New("must supply a number")
	}
	return nil
}

func (m ModuleContext) Client() api.ClientWrapper {
	return m.client
}

func (m ModuleContext) Logger() api.Log {
	return m.log
}