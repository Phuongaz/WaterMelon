package economy

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/modules"
	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/Phuongaz/minecraft-bedrock-server/src/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type GiveBalance struct {
	Target  string
	Balance int64
}

func (g GiveBalance) Run(src cmd.Source, output *cmd.Output) {
	if _, ok := src.(*player.Player); ok {
		p, ok := server.Global().PlayerByName(g.Target)
		if ok {
			id := p.UUID()
			e := modules.EcoEntry()
			balance, _ := e.Balance(id)
			nbalance := balance + uint64(g.Balance)
			e.Set(id, nbalance)
			p.Messagef("Your balance +%v", nbalance)
		} else {
			output.Errorf("Player %v not found", g.Target)
		}
	}
}

func (g GiveBalance) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.Name())
}
