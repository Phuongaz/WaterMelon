package commands

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/Phuongaz/minecraft-bedrock-server/src/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type SetWorldSpawn struct{}

func (s SetWorldSpawn) Run(src cmd.Source, o *cmd.Output) {
	if !s.Allow(src) {
		o.Errorf("You don't have permission to use this command")
	}
	if p, ok := src.(*player.Player); ok {
		s := cube.PosFromVec3(p.Position())
		server.Global().World().SetSpawn(s)
		o.Printf("Set the world spawn point to (%v, %v, %v)", s[0], s[1], s[2])
	} else {
		o.Error("This command must use in game")
	}
}

func (SetWorldSpawn) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.Name())
}
