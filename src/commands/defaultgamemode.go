package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/phuongaz/minecraft-bedrock-server/src/convert"
	"github.com/phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/phuongaz/minecraft-bedrock-server/src/server"
)

type DefaultGameMode struct {
	GameMode string `cmd:"gamemode"`
}

func (d DefaultGameMode) Run(src cmd.Source, o *cmd.Output) {
	if !d.Allow(src) {
		o.Errorf("You don't have permission to use this command")
	}
	mode, err := convert.ParseGameMode(d.GameMode)
	if err != nil {
		o.Error(err)
		return
	}
	server.Global().World().SetDefaultGameMode(mode)
	o.Printf("Set default game mode to %v", convert.MustString(convert.DumpGameMode(mode)))
}

func (d DefaultGameMode) Allow(s cmd.Source) bool {
	if _, ok := s.(*player.Player); ok {
		return permission.OpEntry().Has(s.(*player.Player).Name())
	}
	return true
}
