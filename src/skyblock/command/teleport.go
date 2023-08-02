package command

import (
	"strconv"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/phuongaz/minecraft-bedrock-server/src/skyblock"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

// Teleport implements a /plot tp command which may be used to teleport to a specific plot owned by the
// player.
type Teleport struct {
	Teleport cmd.SubCommand `cmd:"tp"`
	// Number is the number of the plot to teleport to. These numbers may be found by running /p list.
	Number plotNumber `name:"number"`
}

// Run ...
func (t Teleport) Run(source cmd.Source, output *cmd.Output) {
	p := source.(*player.Player)
	h, _ := skyblock.LookupHandler(p)

	plotPositions := h.PlotPositions()

	number, _ := strconv.Atoi(string(t.Number))
	if number < 1 || number > len(plotPositions) {
		output.Errorf("Unknown plot with number %v. Use /p list to get a list of plots to teleport to.", t.Number)
		return
	}
	pl := h.Plots()[number-1]
	pos := plotPositions[number-1]

	p.Teleport(pos.TeleportPosition(h.Settings()))

	f := pl.ColourToFormat()
	output.Printf(text.Colourf("<%v>■</%v> <green>Successfully teleported to your plot.</green>", f, f))
}

// plotNumber ...
type plotNumber cmd.Varargs

// Options returns a number for every plot the player has.
func (plotNumber) Options(source cmd.Source) []string {
	p := source.(*player.Player)
	h, _ := skyblock.LookupHandler(p)
	m := make([]string, len(h.Plots()))
	for i := range h.Plots() {
		m[i] = strconv.Itoa(i + 1)
	}
	return m
}
