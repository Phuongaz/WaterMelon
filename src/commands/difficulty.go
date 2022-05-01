package commands

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/convert"
	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/Phuongaz/minecraft-bedrock-server/src/server"
	"github.com/df-mc/dragonfly/server/cmd"
)

type Difficulty struct {
	Diff string
}

func (d Difficulty) Run(src cmd.Source, o *cmd.Output) {
	if di, err := convert.ParseDifficulty(d.Diff); err != nil {
		o.Error(err)
	} else {
		server.Global().World().SetDifficulty(di)
		o.Printf("Set game difficulty to %v", convert.MustString(convert.DumpDifficulty(di)))
	}
}

func (d Difficulty) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.Name())
}
