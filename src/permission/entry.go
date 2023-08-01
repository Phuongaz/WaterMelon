package permission

import (
	"strings"
	"sync"

	"github.com/phuongaz/minecraft-bedrock-server/src/util"
)

type Entry struct {
	mu     sync.Mutex
	path   string
	list   []string
	except string
}

func (e *Entry) write() {
	util.MustWriteFile(e.path, []byte(strings.Join(e.list, "\n")))
}

func (e *Entry) Reload() {
	e.mu.Lock()
	defer e.mu.Unlock()
	if !util.FileExist(e.path) {
		util.MustWriteFile(e.path, nil)
	}
	var s []string
	for _, a := range strings.Split(string(util.MustReadFile(e.path)), "\n") {
		if len(strings.TrimSpace(a)) != 0 {
			s = append(s, a)
		}
	}
	e.list = s
}

func (e *Entry) GetAll() []string {
	e.mu.Lock()
	defer e.mu.Unlock()
	arr := make([]string, len(e.list))
	copy(arr, e.list)
	return arr
}

func (e *Entry) Has(n string) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.hasNoLock(n)
}

func (e *Entry) hasNoLock(n string) bool {
	if n == e.except {
		return true
	}
	for _, l := range e.list {
		if l == n {
			return true
		}
	}
	return false
}

func (e *Entry) Add(n string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if !e.hasNoLock(n) {
		e.list = append(e.list, n)
		e.write()
	}
}

func (e *Entry) Delete(n string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.hasNoLock(n) {
		var a []string
		for _, l := range e.list {
			if l != n {
				a = append(a, l)
			}
		}
		e.list = a
		e.write()
	}
}

func NewEntry(path string, expect string) *Entry {
	e := &Entry{
		mu:     sync.Mutex{},
		path:   path,
		list:   nil,
		except: expect,
	}
	e.Reload()
	return e
}
