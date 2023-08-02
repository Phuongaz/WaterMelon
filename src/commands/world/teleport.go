package world

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/phuongaz/minecraft-bedrock-server/src/server"
)

type Teleport struct {
	Teleport cmd.SubCommand `cmd:"tp"`
	Name     string         `cmd:"name"`
}

func (t Teleport) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		if t.Allow(src) {
			worldName := t.Name
			world, ok := server.WaterMelonGlobal().WorldManager.GetWorld(worldName)
			if !ok {
				output.Errorf("World %v not found", worldName)
				return
			}
			p.Teleport(world.Spawn().Vec3())
			output.Printf("Teleported to world %v", p.World().Name())
		} else {
			output.Errorf("You don't have permission to use this command")
		}
	}
}

func (t Teleport) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.(*player.Player).Name())
}
