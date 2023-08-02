package skyblock

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/chunk"
)

// Generator implements a generator for a plot world. The settings of the generator are configurable,
// allowing for different results depending on the fields set.
type Generator struct {
	floor, boundary, road uint32
	width                 int
}

// NewGenerator returns a new plot Generator with the Settings passed.
func NewGenerator(s Settings) *Generator {
	return &Generator{
		width: s.PlotWidth,
	}
}

// GenerateChunk generates a chunk for a plot world.
func (g *Generator) GenerateChunk(pos world.ChunkPos, chunk *chunk.Chunk) {
	chunkX := pos.X()
	chunkZ := pos.Z()
	if chunkX%20 == 0 {
		if chunkZ%20 == 0 {
			for x1 := 4; x1 < 11; x1++ {
				for z1 := 4; z1 < 11; z1++ {
					chunk.SetBlock(uint8(x1), 4, uint8(z1), 0, world.BlockRuntimeID(block.Grass{}))
				}
			}
			for x2 := 5; x2 < 10; x2++ {
				for z2 := 5; z2 < 10; z2++ {
					chunk.SetBlock(uint8(x2), 3, uint8(z2), 0, world.BlockRuntimeID(block.Dirt{}))
				}
			}
			for x3 := 6; x3 < 9; x3++ {
				for z3 := 6; z3 < 9; z3++ {
					chunk.SetBlock(uint8(x3), 9, uint8(z3), 0, world.BlockRuntimeID(block.Leaves{}))
					chunk.SetBlock(uint8(x3), 2, uint8(z3), 0, world.BlockRuntimeID(block.Dirt{}))
				}
			}
		}
	}
}
