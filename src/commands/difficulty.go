package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/phuongaz/minecraft-bedrock-server/src/convert"
	"github.com/phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/phuongaz/minecraft-bedrock-server/src/server"
)

type Difficulty struct {
	Diff string
}

func (d Difficulty) Run(src cmd.Source, o *cmd.Output) {
	if !d.Allow(src) {
		o.Errorf("You don't have permission to use this command")
	}
	if di, err := convert.ParseDifficulty(d.Diff); err != nil {
		o.Error(err)
	} else {
		server.Global().World().SetDifficulty(di)
		o.Printf("Set game difficulty to %v", convert.MustString(convert.DumpDifficulty(di)))
	}
}

func (d Difficulty) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.(*player.Player).Name())
}
