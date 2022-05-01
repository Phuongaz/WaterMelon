package commands

import (
	"runtime"

	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"

	"github.com/df-mc/dragonfly/server/cmd"
)

type Status struct{}

func (Status) Run(src cmd.Source, o *cmd.Output) {
	stat := getMemStats()
	o.Printf("Goroutine Count: %v", runtime.NumGoroutine())
	o.Printf("Allocated Memory: %dMB", stat.Sys/1024/1024)
	o.Printf("Virtual Memory: %dMB", stat.HeapSys/1024/1024)
	o.Printf("Stack Memory: %dMB", stat.StackSys/1024/1024)
	o.Printf("Heap Object: %d", (stat.Mallocs-stat.Frees)/1024/1024)
	o.Printf("GC cycles: %d", stat.NumGC)
}

func (Status) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.Name())
}

func getMemStats() runtime.MemStats {
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)
	return m2
}
