package npc

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/npc"
	"github.com/phuongaz/minecraft-bedrock-server/src/permission"
)

type Create struct {
	Sub           create
	Name          string
	CommandExcute string
}

func (d Create) Run(src cmd.Source, _ *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		if d.Allow(src) {
			settings := npc.Settings{
				Name:     d.Name,
				Position: p.Position(),
				Skin:     p.Skin(),
			}
			f := func(player *player.Player) {
				player.ExecuteCommand(d.CommandExcute)
			}
			npc.Create(settings, p.World(), f)
		}
	}
}

type create string

func (create) SubName() string {
	return "create"
}

func (d Create) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.(*player.Player).Name())
}
