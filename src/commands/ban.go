package commands

import (
	"github.com/df-mc/dragonfly/server/player"
	"sort"
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/phuongaz/minecraft-bedrock-server/src/permission"
	"github.com/phuongaz/minecraft-bedrock-server/src/server"
)

type Ban struct {
	Target string
}

func (b Ban) Run(src cmd.Source, o *cmd.Output) {
	defer o.Messages()
	if b.Target == "" {
		o.Error("Command argument error")
		return
	}
	if t, found := server.Global().PlayerByName(b.Target); found {
		t.Disconnect("Banned by admin")
	}
	permission.BanEntry().Add(b.Target)
	o.Printf("Banned player %v", b.Target)
}

func (b Ban) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.(*player.Player).Name())
}

type Unban struct {
	Target string `cmd:"target"`
}

func (u Unban) Run(src cmd.Source, o *cmd.Output) {
	if u.Target == "" {
		o.Error("Command argument error")
		return
	}
	permission.BanEntry().Delete(u.Target)
	o.Printf("Unbanned player %v", u.Target)
}

func (u Unban) Allow(s cmd.Source) bool {
	return permission.OpEntry().Has(s.(*player.Player).Name())
}

type BanList struct {
}

func (BanList) Run(src cmd.Source, o *cmd.Output) {
	arr := permission.BanEntry().GetAll()
	sort.Strings(arr)
	o.Printf("There are %v total banned players:", len(arr))
	o.Print(strings.Join(arr, ", "))
}

func (b BanList) Allow(s cmd.Source) bool {
	if _, ok := s.(*player.Player); ok {
		return permission.OpEntry().Has(s.(*player.Player).Name())
	}
	return true
}
