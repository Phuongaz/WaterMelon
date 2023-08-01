package modules

import (
	"github.com/provsalt/economy"
	"github.com/provsalt/economy/handler"
	"github.com/provsalt/economy/provider"
)

func EcoEntry() (e *economy.Economy) {
	sql, err := provider.NewSQLite("test.sqlite")
	if err != nil {
		panic(err)
	}
	e = economy.New(sql, handler.NopEconomyHandler{})
	return
}
