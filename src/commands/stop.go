package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/phuongaz/minecraft-bedrock-server/src/server"
)

type Stop struct{}

func (s Stop) Run(src cmd.Source, o *cmd.Output) {
	if !s.Allow(src) {
		o.Errorf("You don't have permission to use this command")
	}
	out := &cmd.Output{}
	out.Print("Stopping the server")
	for _, p := range server.Global().Players() {
		p.SendCommandOutput(out)
	}
	_ = server.Global().Close()
}

func (Stop) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.(*player.Player).Name())
}
