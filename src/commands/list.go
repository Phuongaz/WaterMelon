package commands

import (
	"sort"
	"strings"

	"github.com/phuongaz/minecraft-bedrock-server/src/server"
	"github.com/df-mc/dragonfly/server/cmd"
)

type List struct{}

func (List) Run(_ cmd.Source, o *cmd.Output) {
	players := server.Global().Players()
	names := make([]string, len(players))
	for i, p := range players {
		names[i] = p.Name()
	}
	sort.Strings(names)
	o.Printf("There are %v/%v players online:", len(players), server.Global().MaxPlayerCount())
	o.Print(strings.Join(names, ", "))
}
