package app

import (
	"github.com/GreatGodApollo/GoModz/pkg/api"
	"github.com/bwmarrin/discordgo"
	"os"
	"path"
	"plugin"
	"strings"
)

func (mb *ModzBot) HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Author.Bot && mb.IgnoreBots {
		return
	}

	var prefix string
	var contains bool
	var err error
	for i := 0; i < len(mb.Prefixes); i++ {
		prefix = mb.Prefixes[i]
		if strings.HasPrefix(m.Content, prefix) {
			contains = true
			break
		}
	}

	if !contains {
		return
	}

	cmd := strings.Split(strings.TrimPrefix(m.Content, prefix), " ")
	channel, _ := s.Channel(m.ChannelID)

	if command, exist, _ := mb.GetCommand(cmd[0]); exist {
		var inDm bool
		if channel.Type == discordgo.ChannelTypeDM {
			inDm = true
		}

		// Check the user's permissions
		if command.Meta().Type != api.CommandTypePM && !inDm && !api.CheckSessionPermissions(s, m.GuildID, m.Author.ID, command.Meta().UserPermissions) {
			if api.CheckSessionPermissions(s, m.GuildID, s.State.User.ID, api.PermissionMessagesEmbedLinks) {
				embed := api.NewEmbedBuilder().
					SetTitle("Insufficient Permission!").
					SetDescription("You don't have the required permissions to run this command!").
					SetColor(0xff0000).Build()

				if !command.Meta().Hidden {
					_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
				}
			} else {
				if !command.Meta().Hidden {
					_, err = s.ChannelMessageSend(m.ChannelID, ":x: You don't have the required permissions to run this command! :x:")
				}
			}
			if err != nil {
				mb.OnErrorFunc(mb, CommandContext{}, err)
			}
			mb.GetLogger().Debugf("P: FALSE C: %s[%s] U: %s#%s[%s] M: %s", channel.Name, m.ChannelID, m.Author.Username, m.Author.Discriminator, m.Author.ID, m.Content)
			return
		}

		// Check the bot's permissions
		if command.Meta().Type != api.CommandTypePM && !inDm && !api.CheckSessionPermissions(s, m.GuildID, s.State.User.ID, command.Meta().BotPermissions) {
			if api.CheckSessionPermissions(s, m.GuildID, s.State.User.ID, api.PermissionMessagesEmbedLinks) {
				embed := api.NewEmbedBuilder().
					SetTitle("Insufficient Permission!").
					SetDescription("I don't have the required permissions to run this command!").
					SetColor(0xff0000).Build()

				if !command.Meta().Hidden {
					_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
				}
			} else {
				if !command.Meta().Hidden {
					_, err = s.ChannelMessageSend(m.ChannelID, ":x: I don't have the required permissions to run this command! :x:")
				}
			}

			if err != nil {
				mb.OnErrorFunc(mb, CommandContext{}, err)
			}

			mb.GetLogger().Debugf("P: FALSE C: %s[%s] U: %s#%s[%s] M: %s", channel.Name, m.ChannelID, m.Author.Username, m.Author.Discriminator, m.Author.ID, m.Content)
			return
		}

		// Make sure it's being used in the correct place
		if channel.Type == discordgo.ChannelTypeDM && command.Meta().Type == api.CommandTypeGuild {
			if api.CheckSessionPermissions(s, m.GuildID, s.State.User.ID, api.PermissionMessagesEmbedLinks) {
				embed := api.NewEmbedBuilder().
					SetTitle("Invalid Channel!").
					SetDescription("You cannot run this command in a private message.").
					SetColor(0xff0000).Build()

				if !command.Meta().Hidden {
					_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
				}
			} else {
				if !command.Meta().Hidden {
					_, err = s.ChannelMessageSend(m.ChannelID, ":x: You cannot run this command in a private message. :x:")
				}
			}

			if err != nil {
				mb.OnErrorFunc(mb, CommandContext{}, err)
			}
			mb.GetLogger().Debugf("P: FALSE C: %s[%s] U: %s#%s[%s] M: %s", channel.Name, m.ChannelID, m.Author.Username, m.Author.Discriminator, m.Author.ID, m.Content)
			return
		} else if channel.Type == discordgo.ChannelTypeGuildText && command.Meta().Type == api.CommandTypePM {
			if api.CheckSessionPermissions(s, m.GuildID, s.State.User.ID, api.PermissionMessagesEmbedLinks) {
				embed := api.NewEmbedBuilder().
					SetTitle("Invalid Channel!").
					SetDescription("You cannot run this command in a guild.").
					SetColor(0xff0000).Build()

				if !command.Meta().Hidden {
					_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
				}
			} else {
				if !command.Meta().Hidden {
					_, err = s.ChannelMessageSend(m.ChannelID, ":x: You cannot run this command in a guild. :x:")
				}
			}

			if err != nil {
				mb.OnErrorFunc(mb, CommandContext{}, err)
			}
			mb.GetLogger().Debugf("P: FALSE C: %s[%s] U: %s#%s[%s] M: %s", channel.Name, m.ChannelID, m.Author.Username, m.Author.Discriminator, m.Author.ID, m.Content)
			return
		}

		if command.Meta().OwnerOnly && !mb.IsOwner(m.Author.ID) {
			if api.CheckSessionPermissions(s, m.GuildID, s.State.User.ID, api.PermissionMessagesEmbedLinks) {
				embed := api.NewEmbedBuilder().
					SetTitle("You can't run that command!").
					SetDescription("Sorry, only bot owners can run that!").
					SetColor(0xff0000).Build()

				if !command.Meta().Hidden {
					_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
				}
			} else {
				if !command.Meta().Hidden {
					_, err = s.ChannelMessageSend(m.ChannelID, ":x: Only bot owners can run that command. :x:")
				}
			}

			if err != nil {
				mb.OnErrorFunc(mb, CommandContext{}, err)
			}
			mb.GetLogger().Debugf("P: FALSE C: %s[%s] U: %s#%s[%s] M: %s", channel.Name, m.ChannelID, m.Author.Username, m.Author.Discriminator, m.Author.ID, m.Content)
			return
		}

		mb.GetLogger().Debugf("P: TRUE C: %s[%s] U: %s#%s[%s] M: %s", channel.Name, m.ChannelID, m.Author.Username, m.Author.Discriminator, m.Author.ID, m.Content)
		guild, _ := s.Guild(m.GuildID)
		member, _ := s.State.Member(m.GuildID, m.Author.ID)

		cmdCtx := CommandContext{
			client: mb,
			message: m.Message,
			user: m.Author,
			channel: channel,
			guild: guild,
			member: member,
		}

		err := command.Exec(cmdCtx, cmd[1:])
		if err != nil {
			mb.OnErrorFunc(mb, cmdCtx, err)
		}
	}

}

func (mb *ModzBot) AddPrefix(prefix string) {
	mb.Prefixes = append(mb.Prefixes, prefix)
}

func (mb *ModzBot) RemovePrefix(prefix string) {
	for i, v := range mb.Prefixes {
		if v == prefix {
			mb.Prefixes = append(mb.Prefixes[:i], mb.Prefixes[i+1:]...)
			break
		}
	}
}

func (mb *ModzBot) SetPrefixes(prefixes []string) {
	mb.Prefixes = prefixes
}

func (mb *ModzBot) GetPrefixes() []string {
	return mb.Prefixes
}

func (mb *ModzBot) RegisterCommand(cmd api.Command) {
	if _, exists, _ := mb.GetCommand(cmd.Meta().Name); !exists {
		*mb.Commands = append(*mb.Commands, &cmd)
	}
}

func (mb *ModzBot) GetCommand(name string) (cmd api.Command, exists bool, index int) {
	for i := range *mb.Commands {
		c := *(*mb.Commands)[i]
		if c.Meta().Name == name {
			return c, true, i
		}
		for _, a := range c.Meta().Aliases {
			if a == name {
				return c, true, i
			}
		}
	}
	return nil, false, -1
}

func (mb *ModzBot) UnregisterCommand(name string) {
	if _, exists, index := mb.GetCommand(name); exists {
		*mb.Commands = removeCommandFromSlice(*mb.Commands, index)
	}
}

func (mb *ModzBot) IsOwner(id string) bool {
	for _, o := range mb.Owners {
		if id == o {
			return true
		}
	}
	return false
}

func (mb *ModzBot) GetLogger() api.Log {
	return mb.Logger
}

func (mb *ModzBot) GetSession() *discordgo.Session {
	return mb.Session
}

func (mb *ModzBot) UnloadModules() {
	mdCtx := ModuleContext{
		client: mb,
		log: mb.GetLogger(),
	}
	for _, mod := range *mb.Modules {
		m := *mod
		err := m.DeInit(mdCtx)
		if err != nil {
			mb.GetLogger().Errorf("Failed to unload %s module safely: %s", m.Meta().Name, err.Error())
		}
	}
}

func (mb *ModzBot) LoadModules() error {
	if _, err := os.Stat(api.PluginsDir); err != nil {
		return err
	}

	plugins, err := listFiles(api.PluginsDir, `.*_mod.so`)
	if err != nil {
		return err
	}

	for _, potMod := range plugins {
		plug, err := plugin.Open(path.Join(api.PluginsDir, potMod.Name()))
		if err != nil {
			mb.GetLogger().Errorf("Failed to open plugin %s: %v", potMod.Name(), err)
			continue
		}
		modSym, err := plug.Lookup(api.ModSymbolName)
		if err != nil {
			mb.GetLogger().Errorf("Plugin %s does not export symbol \"%s\"",
				potMod.Name(), api.ModSymbolName)
			continue
		}
		mod, ok := modSym.(api.Mod)
		if !ok {
			mb.GetLogger().Errorf("Symbol %s (from %s) does not implement Commands interface",
				api.ModSymbolName, potMod.Name())
			continue
		}
		if err := mb.RegisterModule(mod); err != nil {
			mb.GetLogger().Errorf("%s initialization failed: %v", potMod.Name(), err)
			continue
		}
	}
	return nil
}

func (mb *ModzBot) RegisterModule(mod api.Mod) error {
	mdCtx := ModuleContext{
		client: mb,
		log: mb.GetLogger(),
	}

	err := mod.Init(mdCtx)
	if err != nil {
		return err
	}

	*mb.Modules = append(*mb.Modules, &mod)

	for _, cmd := range mod.Meta().Commands {
		mb.RegisterCommand(cmd)
	}

	for _, evt := range mod.Meta().Events {
		mb.GetLogger().Debugf("Initializing Event: %s", evt.Meta.Name)
		if cb := mb.Session.AddHandler(evt.Exec); cb == nil {
			mb.GetLogger().Errorf("Event '%s' was unable to be initialized.", evt.Meta.Name)
		}
	}

	return nil
}

func (mb *ModzBot) Modz() *[]*api.Mod {
	return mb.Modules
}