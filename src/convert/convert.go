package convert

import (
	"fmt"
	"strings"

	"github.com/df-mc/dragonfly/server/world"
)

func ParseGameMode(v string) (world.GameMode, error) {
	switch strings.ToLower(v) {
	case "0", "s", "survival":
		return world.GameModeSurvival, nil
	case "1", "c", "creative":
		return world.GameModeCreative, nil
	case "2", "a", "adventure":
		return world.GameModeAdventure, nil
	case "3", "spectator":
		return world.GameModeSpectator, nil
	}
	return nil, fmt.Errorf("unknown gamemode %v", v)
}

func DumpGameMode(g world.GameMode) (string, error) {
	switch g {
	case world.GameModeSurvival:
		return "survival", nil
	case world.GameModeCreative:
		return "creative", nil
	case world.GameModeAdventure:
		return "adventure", nil
	case world.GameModeSpectator:
		return "spectator", nil
	default:
		return "", fmt.Errorf("unknown %T", g)
	}
}

func ParseDifficulty(v string) (world.Difficulty, error) {
	switch strings.ToLower(v) {
	case "0", "p", "peaceful":
		return world.DifficultyPeaceful, nil
	case "1", "e", "easy":
		return world.DifficultyEasy, nil
	case "2", "n", "normal":
		return world.DifficultyNormal, nil
	case "3", "h", "hard":
		return world.DifficultyHard, nil
	}
	return nil, fmt.Errorf("unknown difficulty %v", v)
}

func DumpDifficulty(d world.Difficulty) (string, error) {
	switch d {
	case world.DifficultyPeaceful:
		return "peaceful", nil
	case world.DifficultyEasy:
		return "easy", nil
	case world.DifficultyNormal:
		return "normal", nil
	case world.DifficultyHard:
		return "hard", nil
	default:
		return "", fmt.Errorf("unknown difficulty %T", d)
	}
}

func MustString(e string, err error) string {
	if err != nil {
		panic(err)
	}
	return e
}

func MustByteSlice(e []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return e
}
