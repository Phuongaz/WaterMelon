package commands

import (
	"github.com/df-mc/dragonfly/server/player"
	"runtime"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/phuongaz/minecraft-bedrock-server/src/permission"
)

type GC struct{}

func (GC) Run(src cmd.Source, o *cmd.Output) {
	if permission.OpEntry().Has(src.(*player.Player).Name()) {
		a, b := gc()
		o.Printf("Allocated Memory freed: %v MB", (b.Sys-a.Sys)/1024/1024)
	} else {
		o.Error("You are not operator")
	}
}

func (GC) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.(*player.Player).Name())
}

func gc() (runtime.MemStats, runtime.MemStats) {
	var m runtime.MemStats
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m)
	runtime.GC()
	runtime.ReadMemStats(&m2)
	return m, m2
}
