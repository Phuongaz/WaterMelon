package economy

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/phuongaz/minecraft-bedrock-server/src/modules"
	"github.com/phuongaz/minecraft-bedrock-server/src/server"
)

type PayBalance struct {
	Target  string
	Balance int64
}

func (g PayBalance) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		t, ok := server.Global().PlayerByName(g.Target)
		if ok {
			id := t.UUID()
			pid := p.UUID()
			e := modules.EcoEntry()
			tbalance, _ := e.Balance(id)
			pbalance, _ := e.Balance(pid)
			tnbalance := tbalance + uint64(g.Balance)
			pnbalance := pbalance - uint64(g.Balance)
			e.Set(id, tnbalance)
			e.Set(pid, pnbalance)
			t.Messagef("Player %v pay %v", p.Name(), g.Balance)
			p.Messagef("Pay %v, %v ok", t.Name(), g.Balance)
			err := e.Close()
			if err != nil {
				return
			}
		} else {
			output.Errorf("Player %v not found", g.Target)
		}
	}
}
