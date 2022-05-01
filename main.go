package main

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/commands"
	"github.com/Phuongaz/minecraft-bedrock-server/src/console"
	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/Phuongaz/minecraft-bedrock-server/src/server"
	"github.com/Phuongaz/minecraft-bedrock-server/src/util"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	if err := server.Setup(log); err != nil {
		logrus.Fatal(err)
	}

	chat.Global.Subscribe(&util.LoggerSubscriber{Logger: log})
	commands.Setup()
	c := console.Setup(log)
	c.Run()
	server.CloseOnProgramEnd(log, func() {
		c.Stop()
	})
	if err := server.Global().Start(); err != nil {
		logrus.Fatal(err)
	}

	server.Global().Allow(permission.BanEntry().ServerAllower("You are banned", false))
	server.Loop(func(p *player.Player) {
	}, func() {
	})
}
