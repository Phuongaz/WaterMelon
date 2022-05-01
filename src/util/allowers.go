package util

import (
	"net"

	"github.com/df-mc/dragonfly/server"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
)

type serverAllower struct {
	a server.Allower
	b server.Allower
}

func (s *serverAllower) Allow(addr net.Addr, i login.IdentityData, c login.ClientData) (string, bool) {
	if m, b := s.a.Allow(addr, i, c); !b {
		return m, b
	}
	return s.b.Allow(addr, i, c)
}

func LinkServerAllower(a, b server.Allower) server.Allower {
	return &serverAllower{a: a, b: b}
}
