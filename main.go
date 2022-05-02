package main

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/commands"
	"github.com/Phuongaz/minecraft-bedrock-server/src/console"
	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/Phuongaz/minecraft-bedrock-server/src/server"
	"github.com/Phuongaz/minecraft-bedrock-server/src/skyblock"
	"github.com/Phuongaz/minecraft-bedrock-server/src/util"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
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

	w := server.Global().World()
	w.SetDefaultGameMode(world.GameModeCreative)
	w.SetSpawn(cube.Pos{2, skyblock.RoadHeight, 2})
	w.SetTime(5000)
	w.StopTime()
	settings := skyblock.Settings{
		PlotWidth:    32,
		MaximumPlots: 16,
	}
	db, err := skyblock.OpenDB("plots", settings)
	if err != nil {
		log.Fatalf("error opening plot database: %v", err)
	}

	if err != nil {
		log.Fatalf("error opening plot database: %v", err)
	}
	w.Generator(skyblock.NewGenerator(settings))
	w.Handle(skyblock.NewWorldHandler(w, settings))

	for {
		p, err := server.Global().Accept()
		if err != nil {
			break
		}
		p.Handle(skyblock.NewPlayerHandler(p, settings, db))
	}
	_ = db.Close()
}
