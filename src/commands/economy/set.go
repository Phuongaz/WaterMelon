package economy

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/modules"
	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/Phuongaz/minecraft-bedrock-server/src/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type SetBalance struct {
	Target  string
	Balance int64
}

func (b SetBalance) Run(src cmd.Source, output *cmd.Output) {
	if _, ok := src.(*player.Player); ok {
		p, ok := server.Global().PlayerByName(b.Target)
		if ok {
			id := p.UUID()
			e := modules.EcoEntry()
			e.Set(id, uint64(b.Balance))
			p.Messagef("Ok")
			e.Close()
		} else {
			output.Errorf("Player %v not found", b.Target)
		}
	}
}

func (SetBalance) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.Name())
}
