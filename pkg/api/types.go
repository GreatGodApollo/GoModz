package api

import (
	"github.com/bwmarrin/discordgo"
	"io"
)

type Permission int
type CommandType int
type LogLevel uint32

const Version string = "0.1.0"

const (
	CommandTypePM CommandType = iota

	CommandTypeGuild

	CommandTypeEverywhere
)

const (
	LogLevelPanic LogLevel = iota
	LogLevelFatal
	LogLevelError
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
	LogLevelTrace
)

type Module interface {
	Init(ModuleContext) error
	DeInit(ModuleContext) error
}

type Log interface {
	Trace(...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	Panic(...interface{})

	/// Formatted logs
	Tracef(string, ...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
	Panicf(string, ...interface{})

	SetLevel(LogLevel)
	GetLevel() LogLevel
}

type ClientWrapper interface {
	AddPrefix(string)
	RemovePrefix(string)
	SetPrefixes([]string)
	GetPrefixes() []string

	RegisterModule(Mod) error
	Modz() *[]*Mod

	RegisterCommand(Command)
	GetCommand(string) (cmd Command, exists bool, index int)
	UnregisterCommand(string)

	IsOwner(string) bool

	GetLogger() Log

	GetSession() *discordgo.Session
}

type CommandContext interface {
	Client() ClientWrapper
	Message() *discordgo.Message
	User() *discordgo.User
	Channel() *discordgo.Channel
	Guild() *discordgo.Guild
	Member() *discordgo.Member

	Reply(string) (*discordgo.Message, error)
	ReplyEmbed(*discordgo.MessageEmbed) (*discordgo.Message, error)
	ReplyEmbedBuilder(EmbedBuilder) (*discordgo.Message, error)
	ReplyFile(filename string, file io.Reader) (*discordgo.Message, error)
	ReplyTTS(string) (*discordgo.Message, error)

	PurgeMessages(int) error
}

type ModuleContext interface {
	Client() ClientWrapper
	Logger() Log
}

type CommandMeta struct {
	Name string
	Aliases []string
	Description string
	OwnerOnly bool
	Hidden bool
	UserPermissions Permission
	BotPermissions Permission
	Type CommandType
}

type EventMeta struct {
	Name string
	Description string
}

type ModuleMeta struct {
	Name string
	Description string
	Commands []Command
	Events []*Event
}

type Command interface {
	Meta() *CommandMeta
	Exec(context CommandContext, args []string) error
}

type Event struct {
	Meta *EventMeta
	Exec interface{}
}

type Mod interface {
	Meta() *ModuleMeta
	Module
}