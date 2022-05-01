package util

import (
	"fmt"
	"strings"

	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
)

// LoggerSubscriber is an implementation of Subscriber that forwards messages sent to the logger
type LoggerSubscriber struct {
	Logger *logrus.Logger
}

// Message ...
func (c *LoggerSubscriber) Message(a ...interface{}) {
	s := make([]string, len(a))
	for i, b := range a {
		s[i] = fmt.Sprint(b)
	}
	t := text.ANSI(strings.TrimSpace(strings.Join(s, " ")) + "§r")
	c.Logger.Info(t)
}
