package main

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/phuongaz/minecraft-bedrock-server/src/commands"
	"github.com/phuongaz/minecraft-bedrock-server/src/console"
	"github.com/phuongaz/minecraft-bedrock-server/src/server"
	"github.com/phuongaz/minecraft-bedrock-server/src/skyblock"
	"github.com/phuongaz/minecraft-bedrock-server/src/util"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	wm, err := server.New(log)
	if err != nil {
		log.Fatal(err)
	}
	chat.Global.Subscribe(&util.LoggerSubscriber{Logger: log})
	commands.Setup()
	c := console.Setup(log)
	c.Run()
	wm.CloseOnProgramEnd(log, func() {
		c.Stop()
	})

	//wm.Srv.Allow(permission.BanEntry().ServerAllower("You are banned", false))

	settings := skyblock.Settings{
		PlotWidth:    32,
		MaximumPlots: 16,
	}

	w, err := wm.WorldManager.CreateWorld("plots", world.Overworld, skyblock.NewGenerator(settings), cube.Pos{2, skyblock.RoadHeight, 2})
	if err != nil {
		log.Fatalf("error creating plot world: %v", err)
		return
	}

	w.Handle(skyblock.NewWorldHandler(w, settings))
	db, err := skyblock.OpenDB("worlds/plots/db", settings)
	if err != nil {
		log.Fatalf("error opening plot database: %v", err)
	}
	wm.Srv.Listen()
	for wm.Srv.Accept(func(p *player.Player) {
		p.Handle(skyblock.NewPlayerHandler(p, settings, db))
	}) {
	}
	_ = db.Close()
}
