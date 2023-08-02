package economy

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/phuongaz/minecraft-bedrock-server/src/modules"
	"github.com/phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/phuongaz/minecraft-bedrock-server/src/server"
)

type GiveBalance struct {
	Target  string `cmd:"target"`
	Balance int64  `cmd:"balance"`
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
			e.Close()
		} else {
			output.Errorf("Player %v not found", g.Target)
		}
	}
}

func (g GiveBalance) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.(*player.Player).Name())
}
