package command

import (
	"github.com/Phuongaz/minecraft-bedrock-server/src/skyblock"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

// Auto implements the /plot auto command. It teleports the user to the nearest unclaimed plot.
type Auto struct {
	Sub auto
}

// Run ...
func (a Auto) Run(source cmd.Source, output *cmd.Output) {
	p := source.(*player.Player)
	h, _ := skyblock.LookupHandler(p)

	pos := skyblock.PosFromBlockPos(cube.PosFromVec3(p.Position()), h.Settings())

	// We iterate within a growing square, starting at the plots closest to the player and looking up to 16
	// plots around the player in each direction.
	for r := 0; r < 16; r++ {
		for x := -r; x <= r; x++ {
			for z := -r; z <= r; z++ {
				if x == -r || x == r || z == -r || z == r {
					surrounding := pos.Add(skyblock.Position{r, r})
					if _, err := h.DB().Plot(surrounding); err == nil {
						continue
					}
					// The plot isn't yet stored, so it's not claimed. We can teleport the player there.
					p.Teleport(surrounding.TeleportPosition(h.Settings()))
					output.Printf(text.Colourf("<green>A free skyblock was successfully found nearby.</green>"))
					return
				}
			}
		}
	}
	output.Errorf("No free plots could be found in a 32x32 square around you.")
}

// auto ...
type auto string

// SubName ...
func (auto) SubName() string {
	return "auto"
}
