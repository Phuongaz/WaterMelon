package modules

import (
	"fmt"

	"github.com/provsalt/economy"
	"github.com/provsalt/economy/provider"
)

func EcoEntry() economy.Economy {
	sql, err := provider.NewSQLite("test.sqlite")
	if err != nil {
		panic(err)
	}
	e := economy.New(sql)
	if e == nil {
		fmt.Errorf("Economy database notfound")
	}
	return e
}
