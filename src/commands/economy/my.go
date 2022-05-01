package economy

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/modules"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type MyBalance struct{}

func (m MyBalance) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		balance, error := modules.EcoEntry().Balance(p.UUID())
		if error != nil {
			p.Messagef("Your balance %v", balance)
		} else {
			output.Errorf("Error %v", error)
		}
	}
}
