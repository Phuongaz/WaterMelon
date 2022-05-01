package skyblock

// Settings holds the settings for a plot Generator. These settings may be changed in order to change the
// appearance of the plots generated.
type Settings struct {
	// PlotWidth is the width in blocks that each plot generated will be.
	PlotWidth int
	// MaximumPlots is the maximum amount of plots that a player is allowed to claim. Trying to claim more
	// than this will result in an error.
	MaximumPlots int
}
