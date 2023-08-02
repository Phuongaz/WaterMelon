package world

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/phuongaz/minecraft-bedrock-server/src/server"
)

type List struct {
	List string `cmd:"list"`
}

func (l List) Run(src cmd.Source, output *cmd.Output) {
	worlds := server.WaterMelonGlobal().WorldManager.Worlds()
	i := len(worlds)
	output.Printf("There are %v worlds", i)
	for _, world := range worlds {
		time := world.Time()
		output.Printf("World (%v): time: %vms", world.Name(), time)
	}
}

func (l List) Allow(s cmd.Source) bool {
	if _, ok := s.(*player.Player); ok {
		return permission.OpEntry().Has(s.(*player.Player).Name())
	}
	return true
}
