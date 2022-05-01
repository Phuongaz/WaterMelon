package commands

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/convert"
	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/Phuongaz/minecraft-bedrock-server/src/server"
	"github.com/df-mc/dragonfly/server/cmd"
)

type DefaultGameMode struct {
	GameMode string
}

func (d DefaultGameMode) Run(src cmd.Source, o *cmd.Output) {
	mode, err := convert.ParseGameMode(d.GameMode)
	if err != nil {
		o.Error(err)
		return
	}
	server.Global().World().SetDefaultGameMode(mode)
	o.Printf("Set default game mode to %v", convert.MustString(convert.DumpGameMode(mode)))
}

func (d DefaultGameMode) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.Name())
}
