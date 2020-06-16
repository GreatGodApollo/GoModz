package app

import (
	"github.com/GreatGodApollo/GoModz/pkg/api"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type ModzBot struct {
	Session     *discordgo.Session
	config      Configuration
	Prefixes    []string
	Owners      []string
	Logger      api.Log
	Modules     []api.Mod
	Commands    *[]*api.Command
	IgnoreBots  bool
	OnErrorFunc ModzBotOnErrorFunc
}

func NewModzBot(session *discordgo.Session, c Configuration, log api.Log, ignoreBots bool, errorFunc ModzBotOnErrorFunc) ModzBot {
	return ModzBot{
		Session: session,
		config: c,
		Prefixes: c.Prefixes,
		Owners: c.Owners,
		Logger: log,
		IgnoreBots: ignoreBots,
		OnErrorFunc: errorFunc,
		Commands: &[]*api.Command{},
	}
}

type ModzBotOnErrorFunc func(mb *ModzBot, ctx api.CommandContext, err error)

type CommandContext struct {
	client api.ClientWrapper
	message *discordgo.Message
	user *discordgo.User
	channel *discordgo.Channel
	guild *discordgo.Guild
	member *discordgo.Member
}

type ModuleContext struct {
	client api.ClientWrapper
	log api.Log
}

type Logger struct {
	log *logrus.Logger
}

func removeCommandFromSlice(s []*api.Command, i int) []*api.Command {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s) - 1]
}