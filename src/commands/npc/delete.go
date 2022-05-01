package npc

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
)

var _isDelete bool = false

type Delete struct {
	Sub delete
}

func (d Delete) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		if d.Allow(src) {
			_isDelete = true
			p.Message("Tap NPC to delete!")
		}
	}
}

type delete string

func (delete) SubName() string {
	return "delete"
}

func (d Delete) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.Name())
}

type Handler struct {
	player.NopHandler
}

func (h *Handler) HandleHurt(ctx *event.Context, _ *float64, src damage.Source) {
	if src, ok := src.(damage.SourceEntityAttack); ok {
		if attacker, ok := src.Attacker.(*player.Player); ok {
			if _isDelete {
				src.Attacker.Close()
				attacker.Messagef("Delete NPC success!")
				_isDelete = false
			}
		}
	}
}
