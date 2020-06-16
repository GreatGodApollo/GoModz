package app

import (
	"github.com/GreatGodApollo/GoModz/pkg/api"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var Config Configuration
var Log = NewLogger(api.LogLevelDebug)

func Run() {
	Config = LoadConfiguration("config.json", Log)

	Log.Infof("Starting GoModz v%s", api.Version)

	client, err := discordgo.New("Bot " + Config.Token)
	if err != nil {
		Log.Fatal(err.Error())
	}

	mb := NewModzBot(client, Config, Log, true, CommandErrorFunc)

	Log.Info("Registering modules")
	err = mb.LoadModules()
	if err != nil {
		Log.Fatalf("Failed to load modules: %s", err.Error())
	}
	Log.Infof("Loaded %d module(s)", len(mb.Modules))

	client.AddHandler(mb.HandleCommand)
	err = client.Open()
	if err != nil {
		Log.Fatal(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	mb.UnloadModules()
	_ = client.Close()
}

func CommandErrorFunc(mb *ModzBot, ctx api.CommandContext, err error) {
	if (ctx != CommandContext{}) {
		_, err2 := ctx.Reply(":x: Error: `" + err.Error() + "` :x:")
		if err2 != nil {
			mb.GetLogger().Error(err2.Error())
		}
	} else {
		mb.GetLogger().Error(err.Error())
	}
}
