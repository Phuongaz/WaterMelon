package economy

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/phuongaz/minecraft-bedrock-server/src/modules"
)

type MyBalance struct{}

func (m MyBalance) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		e := modules.EcoEntry()
		balance, err2 := e.Balance(p.UUID())
		if err2 != nil {
			p.Messagef("Your balance %v", balance)
		} else {
			output.Errorf("Error %v", err2)
		}
		err := e.Close()
		if err != nil {
			return
		}
	}
}
