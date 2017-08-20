package monitor

import (
	"log"
	"time"
	"github.com/ravichaturvedi/go-monitor/handler"
	"github.com/ravichaturvedi/go-monitor/registry"
	"github.com/ravichaturvedi/go-monitor/scheduler"
	"github.com/ravichaturvedi/go-monitor/server"
	"github.com/ravichaturvedi/go-monitor/plugin"
)


type Monitor interface {
	Start()
}

func New(pluginMap map[string]plugin.Plugin, durationMap map[string]time.Duration) Monitor {
	return defaultMonitor{pluginMap, durationMap}
}


type defaultMonitor struct {
	pluginMap map[string]plugin.Plugin
	durationMap map[string]time.Duration
}

func (m defaultMonitor) Start() {
	r := registry.New(m.pluginMap)
	r = scheduler.New(r, m.durationMap)

	s := server.New(handler.New(r))

	if err := s.Serve(); err != nil {
		log.Fatal("Error while starting server: ", err)
	}
}