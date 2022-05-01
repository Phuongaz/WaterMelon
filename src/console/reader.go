package console

import (
	"bufio"
	"os"
	"strings"
	"sync"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/sirupsen/logrus"
	"go.uber.org/atomic"
)

type Reader struct {
	once sync.Once
	run  atomic.Bool
	sc   *bufio.Scanner
	c    *source
}

func Setup(log *logrus.Logger) *Reader {
	r := &Reader{
		once: sync.Once{},
		run:  atomic.Bool{},
		sc:   bufio.NewScanner(os.Stdin),
		c:    &source{log: log},
	}
	r.run.Store(true)
	return r
}

func (r *Reader) Run() {
	r.once.Do(func() {
		go func() {
			for r.run.Load() {
				r.sc.Scan()
				s := strings.ToValidUTF8(strings.TrimSpace(r.sc.Text()), "")
				if len(s) != 0 {
					args := strings.Split(s, " ")
					name := args[0]
					command, ok := cmd.ByAlias(name)
					if !ok {
						output := &cmd.Output{}
						output.Errorf("Unknown command '%v'", name)
						r.c.SendCommandOutput(output)
						continue
					}
					command.Execute(strings.TrimPrefix(strings.TrimPrefix(s, name), " "), r.c)
				}
			}
		}()
	})
}

func (r *Reader) Stop() {
	r.run.Store(false)
}
