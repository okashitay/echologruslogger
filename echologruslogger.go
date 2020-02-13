package echologruslogger

import (
	"io"
	"github.com/Sirupsen/logrus"
	"os"
	"encoding/json"
	"sync"
	"github.com/labstack/gommon/log"
)

type (
	EchoLogrusLogger struct {
		logger      *logrus.Logger
		level       log.Lvl
		logruslevel logrus.Level
		output      io.Writer
		mutex       sync.Mutex
	}
)

func New(logger *logrus.Logger) (l *EchoLogrusLogger) {
	l = &EchoLogrusLogger{
		logger: logger,
		level:       log.INFO,
		logruslevel: logrus.InfoLevel,
		output:      os.Stdout,
	}
	return
}

func (l *EchoLogrusLogger) SetLevel(v log.Lvl) {
	l.level = v
	var lvl logrus.Level = logrus.PanicLevel
	switch v {
	case log.DEBUG:
		lvl = logrus.DebugLevel
	case log.INFO:
		lvl = logrus.InfoLevel
	case log.WARN:
		lvl = logrus.WarnLevel
	case log.ERROR:
		lvl = logrus.ErrorLevel
	case log.FATAL:
		lvl = logrus.FatalLevel
	}
	l.logruslevel = lvl
	l.logger.Level = lvl
}

func (l *EchoLogrusLogger) SetOutput(w io.Writer) {
	l.output = w
	l.logger.Out = w
}

func (l *EchoLogrusLogger) Print(i ...interface{}) {
	l.logger.Print(i)
}

func (l *EchoLogrusLogger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args)
}

func (l *EchoLogrusLogger) Printj(j log.JSON) {
	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = logrus.GetLevel()
	log.Print(j)
}

func (l *EchoLogrusLogger) Debug(i ...interface{}) {
	l.logger.Debug(i...)
}

func (l *EchoLogrusLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *EchoLogrusLogger) Debugj(j log.JSON) {
	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = l.logruslevel
	log.Out = l.output
	byte, _ := json.Marshal(j)
	log.Debug(string(byte))
}

func (l *EchoLogrusLogger) Info(i ...interface{}) {
	l.logger.Info(i...)
}

func (l *EchoLogrusLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *EchoLogrusLogger) Infoj(j log.JSON) {
	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = logrus.GetLevel()
	log.Out = l.output
	byte, _ := json.Marshal(j)
	log.Info(string(byte))
}

func (l *EchoLogrusLogger) Warn(i ...interface{}) {
	l.logger.Warn(i...)
}

func (l *EchoLogrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *EchoLogrusLogger) Warnj(j log.JSON) {
	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = logrus.GetLevel()
	log.Out = l.output
	byte, _ := json.Marshal(j)
	log.Warn(string(byte))
}

func (l *EchoLogrusLogger) Error(i ...interface{}) {
	l.logger.Error(i...)
}

func (l *EchoLogrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *EchoLogrusLogger) Errorj(j log.JSON) {
	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = logrus.GetLevel()
	log.Out = l.output
	byte, _ := json.Marshal(j)
	log.Error(string(byte))
}

func (l *EchoLogrusLogger) Fatal(i ...interface{}) {
	l.logger.Fatal(i...)
}

func (l *EchoLogrusLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *EchoLogrusLogger) Fatalj(j log.JSON) {
	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Level = logrus.GetLevel()
	log.Out = l.output
	byte, _ := json.Marshal(j)
	log.Fatal(string(byte))
}
