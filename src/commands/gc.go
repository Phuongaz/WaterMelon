package commands

import (
	"runtime"

	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/df-mc/dragonfly/server/cmd"
)

type GC struct{}

func (GC) Run(src cmd.Source, o *cmd.Output) {
	if permission.OpEntry().Has(src.Name()) {
		a, b := gc()
		o.Printf("Allocated Memory freed: %v MB", (b.Sys-a.Sys)/1024/1024)
	} else {
		o.Error("You are not operator")
	}
}

func (GC) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.Name())
}

func gc() (runtime.MemStats, runtime.MemStats) {
	var m runtime.MemStats
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m)
	runtime.GC()
	runtime.ReadMemStats(&m2)
	return m, m2
}
