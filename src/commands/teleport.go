package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/phuongaz/minecraft-bedrock-server/src/permission"
)

type Teleport struct {
	X, Y, Z float64
}

func (t Teleport) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		if ok := t.Allow(src); ok {
			p.Teleport(mgl64.Vec3{t.X, t.Y, t.Z})
			output.Printf("Teleported to X: %d Y: %d Z: %d", int(t.X), int(t.Y), int(t.Z))
		} else {
			output.Errorf("You don't have permission to use this command")
		}
	} else {
		output.Error("This command must use in game")
	}
}

func (Teleport) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.(*player.Player).Name())
}
