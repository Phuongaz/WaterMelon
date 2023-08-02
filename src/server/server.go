package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/df-mc/dragonfly/server"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

var _global *server.Server
var _watermelon *WaterMelon

type WaterMelon struct {
	Log          *logrus.Logger
	Config       server.UserConfig
	Srv          *server.Server
	WorldManager *WorldManager
}

func WaterMelonGlobal() *WaterMelon {
	return _watermelon
}

func Global() *server.Server {
	return _global
}

func New(l *logrus.Logger) (*WaterMelon, error) {
	conf, err := readConfig(l)
	if err != nil {
		return nil, fmt.Errorf("error reading config: %v", err)
	}
	_global = conf.New()
	wm := &WaterMelon{
		Log: l,
		Srv: _global,
	}
	l.Info("Loading World Manager.")
	worldManager := wm.SetupWorldManager()
	wm.WorldManager = worldManager
	_watermelon = wm
	l.Info("Loading worlds...")
	err = wm.WorldManager.LoadAll()
	if err != nil {
		return nil, err
	}
	return wm, nil
}

func readConfig(log server.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return c.Config(log)
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}
	return c.Config(log)
}

func (wm *WaterMelon) CloseOnProgramEnd(log *logrus.Logger, f func()) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func(fn func()) {
		<-c
		if err := wm.WorldManager.CloseAll(); err != nil {
			log.Errorf("error closing worlds: %v", err)
		}
		if err := wm.Srv.Close(); err != nil {
			log.Errorf("error shutting down server: %v", err)
		}
		fn()
	}(f)
}
