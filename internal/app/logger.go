package app

import (
	"github.com/GreatGodApollo/GoModz/pkg/api"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func NewLogger(level api.LogLevel) *Logger {
	grus := logrus.New()
	grus.SetLevel(logrus.Level(level))
	lg := &Logger{
		log: grus,
	}
	grus.Formatter = new(prefixed.TextFormatter)

	return lg
}

func (l *Logger) Trace(i ...interface{}) {
	l.log.Trace(i...)
}

func (l *Logger) Debug(i ...interface{}) {
	l.log.Debug(i...)
}

func (l *Logger) Info(i ...interface{}) {
	l.log.Info(i...)
}

func (l *Logger) Warn(i ...interface{}) {
	l.log.Warn(i...)
}

func(l *Logger) Error(i ...interface{}) {
	l.log.Error(i...)
}

func(l *Logger) Fatal(i ...interface{}) {
	l.log.Fatal(i...)
}

func(l *Logger) Panic(i ...interface{}) {
	l.log.Panic(i...)
}

func(l *Logger) Tracef(s string, i ...interface{}) {
	l.log.Tracef(s, i...)
}

func(l *Logger) Debugf(s string, i ...interface{}) {
	l.log.Debugf(s, i...)
}

func(l *Logger) Infof(s string, i ...interface{}) {
	l.log.Infof(s, i...)
}

func(l *Logger) Warnf(s string, i ...interface{}) {
	l.log.Warnf(s, i...)
}

func(l *Logger) Errorf(s string, i ...interface{}) {
	l.log.Errorf(s, i...)
}

func(l *Logger) Fatalf(s string, i ...interface{}) {
	l.log.Fatalf(s, i...)
}

func(l *Logger) Panicf(s string, i ...interface{}) {
	l.log.Panicf(s, i...)
}

func(l *Logger) SetLevel(level api.LogLevel) {
	l.log.SetLevel(logrus.Level(level))
}

func(l *Logger) GetLevel() api.LogLevel {
	return api.LogLevel(l.log.GetLevel())
}
