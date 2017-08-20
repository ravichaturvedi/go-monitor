package scheduler

import (
	"github.com/ravichaturvedi/go-monitor/registry"
	"time"
)

// Scheduler schedule the plugins to run after every provided duration and exposes the Registry interface.
// This make it a in-place replacement of Registry in case we want to use scheduled plugins and can use the cached data.
//
// If the scheduling information is not found for the plugin name then delegate the call to underlying registry.
type Scheduler interface {
	registry.Registry
}


func New(r registry.Registry, m map[string] time.Duration) {

}


type durationScheduler struct {
	r registry.Registry
	results map[string] resultHolder
}

func (s durationScheduler) Run(pluginName string) (interface{}, error) {
	return s.r.Run(pluginName)
}

func (s durationScheduler) PluginNames() []string {
	return s.r.PluginNames()
}


type resultHolder struct {
	V interface{}
	E error
}