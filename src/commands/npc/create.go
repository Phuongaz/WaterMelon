package npc

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/npc"
)

type Create struct {
	Sub           create
	Name          string
	CommandExcute string
}

func (c Create) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		if c.Allow(src) {
			settings := npc.Settings{
				Name:     c.Name,
				Position: p.Position(),
				Skin:     p.Skin(),
			}
			f := func(player *player.Player) {
				player.ExecuteCommand(c.CommandExcute)
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
	return permission.OpEntry().Has(s.Name())
}
