package npc

import (
	"github.com/df-mc/dragonfly/server/entity"
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/phuongaz/minecraft-bedrock-server/src/permission"
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
	return permission.OpEntry().Has(s.(*player.Player).Name())
}

type Handler struct {
	player.NopHandler
}

func (h *Handler) HandleHurt(ctx *event.Context, damage *float64, attackImmunity *time.Duration, src world.DamageSource) {
	if src, ok := src.(entity.AttackDamageSource); ok {
		if attacker, ok := src.Attacker.(*player.Player); ok {
			if _isDelete {
				err := src.Attacker.Close()
				if err != nil {
					return
				}
				attacker.Messagef("Delete NPC success!")
				_isDelete = false
			}
		}
	}
}
