package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/mcdb"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

var _global *server.Server
var _watermelon *WaterMelon

type WaterMelon struct {
	Log    *logrus.Logger
	Worlds map[string]*world.World
	Config server.UserConfig
	Srv    *server.Server
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
	_ = l
	_global = conf.New()
	wm := &WaterMelon{
		Log:    l,
		Worlds: make(map[string]*world.World),
		Srv:    _global,
	}
	_watermelon = wm
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
		if err := wm.Srv.Close(); err != nil {
			log.Errorf("error shutting down server: %v", err)
		}
		fn()
	}(f)
}

type WorldAlreadyLoadedError struct {
	folder string
}

func (e WorldAlreadyLoadedError) Error() string {
	return fmt.Sprintf("world '%v' is already loaded", e.folder)
}

func (wm *WaterMelon) LoadWorldFromFolder(folder string, dimension world.Dimension, generator world.Generator) (*world.World, error) {
	_, exists := wm.Worlds[folder]
	if exists {
		return nil, WorldAlreadyLoadedError{folder}
	}
	p, err := mcdb.Open(folder)
	if err != nil {
		return nil, fmt.Errorf("error opening world: %v", err)
	}
	w := world.Config{
		Provider:  p,
		Generator: generator,
	}.New()

	wm.Log.Infof("Loaded world '%v'", w.Name())
	settings := p.Settings()
	settings.Name = w.Name()
	p.SaveSettings(settings)
	wm.Worlds[folder] = w
	return w, nil
}

func (wm *WaterMelon) GetWorld(name string) (*world.World, error) {
	w, exists := wm.Worlds[name]
	if !exists {
		return nil, fmt.Errorf("world '%v' not found", name)
	}
	return w, nil
}

func (wm *WaterMelon) CreateWorld(name string, folder string, dimension world.Dimension, generator world.Generator) (*world.World, error) {
	_, exists := wm.Worlds[folder]
	if exists {
		return nil, WorldAlreadyLoadedError{folder}
	}
	p, err := mcdb.Config{}.Open(folder)
	if err != nil {
		return nil, fmt.Errorf("error opening world: %v", err)
	}
	w := world.Config{
		Provider:  p,
		Generator: generator,
	}.New()
	wm.Log.Infof("Created world '%v'", name)
	settings := p.Settings()
	settings.Name = name
	p.SaveSettings(settings)
	wm.Worlds[folder] = w
	return w, nil
}
