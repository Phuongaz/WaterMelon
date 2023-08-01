package economy

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/phuongaz/minecraft-bedrock-server/src/modules"
	"github.com/phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/phuongaz/minecraft-bedrock-server/src/server"
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
			err := e.Set(id, uint64(b.Balance))
			if err != nil {
				return
			}
			p.Messagef("Ok")
			err = e.Close()
			if err != nil {
				return
			}
		} else {
			output.Errorf("Player %v not found", b.Target)
		}
	}
}

func (SetBalance) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.(*player.Player).Name())
}
