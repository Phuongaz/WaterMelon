package commands

import (
	"runtime"

	"github.com/Phuongaz/minecraft-bedrock-server/src/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

type Version struct{}

func (Version) Run(src cmd.Source, o *cmd.Output) {
	o.Printf("This server is running %v", "dragonfly")
	o.Printf("Server version: %v", server.Version())
	o.Printf("Compatible Minecraft version: %v (protocol version: %v)", protocol.CurrentVersion, protocol.CurrentProtocol)
	o.Printf("Golang version: %v", runtime.Version())
	o.Printf("Compiler: %v", runtime.Compiler)
	o.Printf("ARCH/GOODS: %v/%v", runtime.GOARCH, runtime.GOOS)
}
