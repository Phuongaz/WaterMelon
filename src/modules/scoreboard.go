package modules

import (
	"strconv"
	"time"

	"github.com/Phuongaz/minecraft-bedrock-server/src/server"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/scoreboard"
)

func score(p *player.Player) {
	score := scoreboard.New("§l§bWater§cMelon")
	balance, _ := EcoEntry().Balance(p.UUID())
	score.RemovePadding()
	score.Set(0, "§l§bName:§e "+p.Name())
	score.Set(1, "§l§bOnlines:§e "+strconv.FormatInt(int64(server.Global().PlayerCount()), 10))
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
