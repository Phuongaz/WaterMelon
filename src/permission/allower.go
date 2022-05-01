package permission

import (
	"net"

	"github.com/df-mc/dragonfly/server"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
)

func (e *Entry) ServerAllower(msg string, allowHas bool) server.Allower {
	return entryServerAllower{
		e:         e,
		detectHas: allowHas,
		msg:       msg,
	}
}

type entryServerAllower struct {
	e         *Entry
	detectHas bool
	msg       string
}

func (e entryServerAllower) Allow(_ net.Addr, d login.IdentityData, _ login.ClientData) (string, bool) {
	if e.detectHas {
		return e.msg, e.e.Has(d.DisplayName)
	}
	return e.msg, !e.e.Has(d.DisplayName)
}
