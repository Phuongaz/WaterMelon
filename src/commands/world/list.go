package world

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/phuongaz/minecraft-bedrock-server/src/permission"
)

type List struct {
	List string `cmd:"list"`
}

func (l List) Run(src cmd.Source, output *cmd.Output) {
	//TODO: show list worlds
}

func (l List) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.(*player.Player).Name())
}
