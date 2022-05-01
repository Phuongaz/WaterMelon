package server

import (
	"runtime/debug"
)

var _dragonflyVersion = getVersion()

func getVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			// TODO replace support
			if dep.Path == "github.com/df-mc/dragonfly" {
				return dep.Version
			}
		}
	}
	return "Unknown"
}

func Version() string {
	return _dragonflyVersion
}
