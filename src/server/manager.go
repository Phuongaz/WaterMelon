package server

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/mcdb"
	"github.com/df-mc/goleveldb/leveldb/opt"
	_ "github.com/df-mc/goleveldb/leveldb/opt"
	"github.com/sirupsen/logrus"
	"sync"
)

const worldsDir = "worlds/"

type WorldManager struct {
	wm  *WaterMelon
	log *logrus.Logger

	worldsMu  sync.RWMutex
	allWorlds []string
	worlds    map[string]*world.World
}

func (wm *WaterMelon) SetupWorldManager() *WorldManager {
	worldManager := &WorldManager{
		wm:     wm,
		log:    wm.Log,
		worlds: make(map[string]*world.World),
	}
	err := worldManager.LoadAll()
	if err != nil {
		return nil
	}
	return worldManager
}

func (manager *WorldManager) LoadWorld(worldName string) error {
	if _, ok := manager.GetWorld(worldName); ok {
		return fmt.Errorf("world %v is already loaded", worldName)
	}

	manager.log.Debugf("Loading world...")
	p, err := mcdb.Open(worldsDir + worldName)
	if err != nil {
		return fmt.Errorf("error loading world: %v", err)
	}

	p.SaveSettings(&world.Settings{
		Name:  worldName,
		Spawn: cube.Pos{0, -55, 0},
	})

	w := world.Config{
		Dim:      world.Overworld,
		Log:      manager.log,
		ReadOnly: true,
		Provider: p,
	}.New()

	w.SetTickRange(0)
	w.SetTime(6000)
	w.StopTime()

	w.StopWeatherCycle()
	w.SetDefaultGameMode(world.GameModeSurvival)

	manager.worldsMu.Lock()
	manager.worlds[worldName] = w
	manager.worldsMu.Unlock()

	manager.log.Infof(`Loaded world "%v".`, w.Name())
	return nil
}

func (manager *WorldManager) UnloadWorld(w *world.World) error {
	if w == manager.wm.Srv.World() {
		return fmt.Errorf("the default world cannot be unloaded")
	}

	if _, ok := manager.GetWorld(w.Name()); !ok {
		return fmt.Errorf("world isn't loaded")
	}

	manager.log.Debugf("Unloading world '%v'\n", w.Name())
	for _, p := range manager.wm.Srv.Players() {
		if p.World() == w {
			// Teleport all entities from the world, to the default world
			manager.wm.Srv.World().AddEntity(p)
			// Teleport them to the spawn of the world
			p.Teleport(manager.wm.Srv.World().Spawn().Vec3Middle())
		}
	}

	manager.worldsMu.Lock()
	delete(manager.worlds, w.Name())
	manager.worldsMu.Unlock()

	if err := w.Close(); err != nil {
		return fmt.Errorf("error closing world: %v", err)
	}
	manager.log.Debugf("Unloaded world '%v'\n", w.Name())
	return nil
}

func (manager *WorldManager) GetWorld(name string) (*world.World, bool) {
	manager.worldsMu.RLock()
	w, ok := manager.worlds[name]
	manager.worldsMu.RUnlock()
	return w, ok
}

func (manager *WorldManager) CreateWorld(name string, dim world.Dimension, generator world.Generator, pos cube.Pos) (*world.World, error) {
	if _, ok := manager.GetWorld(name); ok {
		return nil, fmt.Errorf("world %v already exists", name)
	}

	conf := mcdb.Config{
		Log:         manager.log,
		Compression: opt.DefaultCompression,
		ReadOnly:    false,
		Entities:    entity.DefaultRegistry,
	}

	open, err := conf.Open(worldsDir + name)
	if err != nil {
		return nil, err
	}

	open.SaveSettings(&world.Settings{
		Name:  name,
		Spawn: pos,
	})

	w := world.Config{
		Dim:       dim,
		Log:       manager.log,
		Provider:  open,
		Generator: generator,
	}.New()

	w.SetTickRange(0)
	w.SetTime(6000)
	w.StopTime()

	w.StopWeatherCycle()
	w.SetDefaultGameMode(world.GameModeSurvival)

	manager.worldsMu.Lock()
	manager.worlds[name] = w
	manager.worldsMu.Unlock()

	manager.log.Infof(`Created world "%v".`, w.Name())
	_ = open.Close()
	return w, nil
}

func (manager *WorldManager) Worlds() []*world.World {
	manager.worldsMu.RLock()
	defer manager.worldsMu.RUnlock()

	var worlds []*world.World
	for _, w := range manager.worlds {
		worlds = append(worlds, w)
	}
	return worlds
}

func (manager *WorldManager) LoadAll() error {
	for _, w := range manager.Worlds() {
		if err := manager.LoadWorld(w.Name()); err != nil {
			return fmt.Errorf("error loading world: %v", err)
		} else {
			manager.log.Infof("Loaded world '%v'", w.Name())
		}
	}
	return nil
}

func (manager *WorldManager) CloseAll() error {
	for _, w := range manager.Worlds() {
		if err := w.Close(); err != nil {
			return fmt.Errorf("error closing world: %v", err)
		} else {
			manager.log.Debugf("Closed world '%v'", w.Name())
		}
	}
	return nil
}
