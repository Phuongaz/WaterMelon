package world

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/Phuongaz/minecraft-bedrock-server/src/server"
	"github.com/df-mc/dragonfly/server/cmd"
)

type List struct {
	Sub list
}

func (l List) Run(src cmd.Source, output *cmd.Output) {
	worldmanager := server.WorldManager()
	list := worldmanager.Worlds()
	output.Printf("Worlds: (%v)", len(list))
	for _, world := range list {
		output.Printf("+ %v", world.Name())
	}
}

func (t List) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.Name())
}

type list string

// SubName ...
func (list) SubName() string {
	return "list"
}
