package commands

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/Phuongaz/minecraft-bedrock-server/src/server"
	"github.com/df-mc/dragonfly/server/cmd"
)

type Stop struct{}

func (Stop) Run(src cmd.Source, o *cmd.Output) {
	out := &cmd.Output{}
	out.Print("Stopping the server")
	for _, p := range server.Global().Players() {
		p.SendCommandOutput(out)
	}
	_ = server.Global().Close()
}

func (Stop) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.Name())
}
