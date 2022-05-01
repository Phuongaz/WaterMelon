package commands

import (
	"fmt"

	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Kick struct {
	Target []cmd.Target
	Reason string `optional:""`
}

func (b Kick) Run(src cmd.Source, o *cmd.Output) {
	if b.Target == nil {
		o.Error("Target not found")
		return
	}
	if len(b.Target) != 1 {
		o.Errorf("Target count cannot be %v", len(b.Target))
		return
	}
	if p, ok := b.Target[0].(*player.Player); ok {
		reason := "Kicked by admin"
		if len(b.Reason) != 0 {
			reason += fmt.Sprintf(": %v", b.Reason)
		}
		p.Disconnect(reason)
		o.Printf("Kicked player %v", b.Target)
	} else {
		o.Error("Target is not a player")
	}
}

func (Kick) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.Name())
}
