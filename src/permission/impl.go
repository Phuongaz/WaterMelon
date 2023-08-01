package permission

import (
	"path/filepath"

	"github.com/phuongaz/minecraft-bedrock-server/src/util"
)

var _banEntry = NewEntry(filepath.Join(util.WorkingPath, "banned-players.txt"), "CONSOLE")
var _opEntry = NewEntry(filepath.Join(util.WorkingPath, "ops.txt"), "CONSOLE")

func BanEntry() *Entry {
	return _banEntry
}

func OpEntry() *Entry {
	return _opEntry
}
