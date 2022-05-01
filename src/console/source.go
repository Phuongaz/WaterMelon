package console

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
)

type source struct {
	log *logrus.Logger
}

func (src source) Name() string {
	return "CONSOLE"
}

func (src source) Position() mgl64.Vec3 {
	return mgl64.Vec3{}
}

func (src source) SendCommandOutput(o *cmd.Output) {
	for _, s := range o.Messages() {
		src.log.Info(text.ANSI(s))
	}
	for _, s := range o.Errors() {
		src.log.Error(text.ANSI(s))
	}
}

func (src source) World() *world.World {
	return nil
}
