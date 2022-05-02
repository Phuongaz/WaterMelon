package world

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/Phuongaz/minecraft-bedrock-server/src/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Teleport struct {
	Sub  tp
	Name string
}

func (t Teleport) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		if t.Allow(src) {
			world_name := t.Name
			world, ok := server.WorldManager().World(world_name)
			if ok {
				p.Teleport(world.Spawn().Vec3())
				p.Messagef("Teleport to %v", world_name)
			} else {
				p.Messagef("World %v not found", world_name)
			}
		} else {
			output.Errorf("You don't have permission to use this command")
		}
	}
}

func (t Teleport) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.Name())
}

type tp string

// SubName ...
func (tp) SubName() string {
	return "tp"
}
