package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/phuongaz/minecraft-bedrock-server/src/convert"
	"github.com/phuongaz/minecraft-bedrock-server/src/permission"
)

type GameMode struct {
	GameMode string
}

func (g GameMode) Run(src cmd.Source, o *cmd.Output) {
	if permission.OpEntry().Has(src.(*player.Player).Name()) {
		if p, ok := src.(*player.Player); ok {
			mode, err := convert.ParseGameMode(g.GameMode)
			if err != nil {
				o.Error(err)
				return
			}
			p.SetGameMode(mode)
			o.Printf("Set game mode to %v", convert.MustString(convert.DumpGameMode(mode)))
		} else {
			o.Error("This command must use in game")
		}
	} else {
		o.Error("You are not operator")
	}
}

func (g GameMode) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.(*player.Player).Name())
}
