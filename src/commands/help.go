package commands

import (
	"fmt"
	"sort"

	"github.com/df-mc/dragonfly/server/cmd"
)

type Help struct{}

func (Help) Run(_ cmd.Source, o *cmd.Output) {
	o.Print("--- Showing help ---")
	cmds := cmd.Commands()
	var a []string
	m := make(map[string]string, len(cmds))
	for _, c := range cmds {
		if _, ok := m[c.Name()]; !ok {
			a = append(a, c.Name())
			m[c.Name()] = fmt.Sprintf("/%v: %v", c.Name(), c.Description())
		}
	}
	sort.Strings(a)
	for _, c := range a {
		o.Printf(m[c])
	}
	o.Print("Tip: Use the <tab> key while typing a command to auto-complete the command or its arguments")
}
