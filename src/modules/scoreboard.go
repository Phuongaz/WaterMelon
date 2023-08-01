package modules

import (
	"strconv"
	"time"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/scoreboard"
	"github.com/phuongaz/minecraft-bedrock-server/src/server"
)

func score(p *player.Player) {
	score := scoreboard.New("§l§bWater§cMelon")
	balance, _ := EcoEntry().Balance(p.UUID())
	score.RemovePadding()
	score.Set(0, "§l§bName:§e "+p.Name())
	score.Set(1, "§l§bOnlines:§e "+strconv.FormatInt(int64(len(server.Global().Players())), 10))
	score.Set(2, "§l§bMoney:§e "+strconv.FormatInt(int64(balance), 10))
	p.SendScoreboard(score)
}

func SendScore(p *player.Player) {
	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		if p == nil {
			ticker.Stop()
			return
		}
		score(p)
	}
}
