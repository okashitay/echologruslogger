package echologruslogger

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/gommon/log"
	. "github.com/smartystreets/goconvey/convey"
)

func test(l *EchoLogrusLogger, t *testing.T) {
	b := new(bytes.Buffer)
	l.SetOutput(b)
	l.SetLevel(log.WARN)

	l.Print("print")
	l.Printf("print%s", "f")
	l.Debug("debug")
	l.Debugf("debug%s", "f")
	l.Info("info")
	l.Infof("info%s", "f")
	l.Warn("warn")
	l.Warnf("warn%s", "f")
	l.Error("error")
	l.Errorf("error%s", "f")

	Convey("assert log text", t, func() {
		str := b.String()
		fmt.Println(str)
		So(str, ShouldNotContainSubstring, "print")
		So(str, ShouldNotContainSubstring, "printf")
		So(str, ShouldNotContainSubstring, "debug")
		So(str, ShouldNotContainSubstring, "debugf")
		So(str, ShouldNotContainSubstring, "info")
		So(str, ShouldNotContainSubstring, "infof")
		So(str, ShouldContainSubstring, `WARN`)
		So(str, ShouldContainSubstring, `warn`)
		So(str, ShouldContainSubstring, `warnf`)
		So(str, ShouldContainSubstring, `ERRO`)
		So(str, ShouldContainSubstring, `errorf`)
	})
}

func TestLog(t *testing.T) {
	l := New(logrus.New())
	test(l, t)
}

func TestFatal(t *testing.T) {
	l := New(logrus.New())
	switch os.Getenv("TEST_LOGGER_FATAL") {
	case "fatal":
		l.Fatal("fatal")
		return
	case "fatalf":
		l.Fatalf("fatal-%s", "f")
		return
	}

	loggerFatalTest(t, "fatal", "fatal")
	loggerFatalTest(t, "fatalf", "fatal-f")
}

func loggerFatalTest(t *testing.T, env string, contains string) {
	Convey("testing fatal", t, func() {
		buf := new(bytes.Buffer)
		cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
		cmd.Env = append(os.Environ(), "TEST_LOGGER_FATAL="+env)
		cmd.Stdout = buf
		cmd.Stderr = buf
		err := cmd.Run()
		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			So(buf.String(), ShouldContainSubstring, contains)
			return
		}
		t.Fatalf("process ran with err %v, want exit status 1", err)
	})
}

func TestFormat(t *testing.T) {
	Convey("testing format", t, func() {
		l := New(logrus.New())
		b := new(bytes.Buffer)
		l.SetLevel(log.INFO)
		l.SetOutput(b)
		l.Info("info")
		So(b.String(), ShouldContainSubstring, "INFO")
		So(b.String(), ShouldContainSubstring, "info")
	})
}

func TestJSON(t *testing.T) {
	Convey("test json", t, func() {
		l := New(logrus.New())
		b := new(bytes.Buffer)
		l.SetOutput(b)
		l.SetLevel(log.DEBUG)
		l.Debugj(log.JSON{"name": "value"})
		So(b.String(), ShouldContainSubstring, `"{\"name\":\"value\"}"`)
	})
}
