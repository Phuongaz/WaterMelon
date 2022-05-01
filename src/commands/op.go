package commands

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/Phuongaz/minecraft-bedrock-server/src/server"
	"github.com/df-mc/dragonfly/server/cmd"
)

type Op struct {
	Target string
}

func (b Op) Run(src cmd.Source, o *cmd.Output) {
	if b.Target == "" {
		o.Error("Command argument error")
		return
	}
	if t, found := server.Global().PlayerByName(b.Target); found {
		op := &cmd.Output{}
		op.Print("You have been opped")
		t.SendCommandOutput(op)
	}
	permission.OpEntry().Add(b.Target)
	o.Printf("Opped: %v", b.Target)
}

func (Op) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.Name())
}

type DeOp struct {
	Target string
}

func (b DeOp) Run(src cmd.Source, o *cmd.Output) {
	if b.Target == "" {
		o.Error("Command argument error")
		return
	}
	permission.OpEntry().Delete(b.Target)
	o.Printf("De-opped: %v", b.Target)
}

func (DeOp) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.Name())
}
