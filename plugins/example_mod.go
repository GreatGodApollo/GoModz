package main

import (
	"github.com/GreatGodApollo/GoModz/internal/app"
	"github.com/GreatGodApollo/GoModz/pkg/api"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// Single command
type exampleCmd struct{}



// Command metadata
func (e exampleCmd) Meta() *api.CommandMeta {
	return &api.CommandMeta{
		Name: "example",
		Aliases: []string{},
		Description: "An example command.",
		OwnerOnly: false,
		Hidden: false,
		UserPermissions: 0,
		BotPermissions: api.PermissionMessagesSend | api.PermissionMessagesEmbedLinks,
		Type: api.CommandTypeEverywhere,
	}
}

// Called when the bot gets the command
func (e exampleCmd) Exec(ctx api.CommandContext, args []string) error {
	ctx.Reply("This is just an example...")
	return nil
}

// Single Event
var exampleEvt = &api.Event{
	Meta: &api.EventMeta{
		Name:        "ExampleEvent",
		Description: "An example event..",
	},
	Exec: func(session *discordgo.Session, event *discordgo.MessageCreate) {
		if !event.Author.Bot && strings.Contains(event.Content, "rip") {
			session.MessageReactionAdd(event.ChannelID, event.Message.ID, app.RegionalF)
		}
	},
}

// Whole module
type exampleMod struct{}

func (t *exampleMod) Init(ctx api.ModuleContext) error {
	ctx.Logger().Info("Example mod init call")
	return nil
}

func (t *exampleMod) DeInit(ctx api.ModuleContext) error {
	ctx.Logger().Info("Example mod deinit call")
	return nil
}

func (t *exampleMod) Meta() *api.ModuleMeta {
	return &api.ModuleMeta{
		Name:     "ExampleMod",
		Description: "An example module.",
		Commands: []api.Command{
			exampleCmd{},
		},
		Events: []*api.Event{
			exampleEvt,
		},
	}
}

// Export it
var Mod exampleMod
