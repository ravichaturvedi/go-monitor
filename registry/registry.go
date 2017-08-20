package registry

import (
	"log"
	"errors"
	"github.com/ravichaturvedi/go-monitor/plugin"
)

var ErrPluginNotFound = errors.New("Plugin with the provided name is not available in registry.")


// Registry holds the available plugins and expose operations to be performed on those plugins.
type Registry interface {

	// Run the specified plugin name
	Run(pluginName string) (interface{}, error)

	// List the available plugin names
	PluginNames() []string
}


func New(plugins ...plugin.Plugin) Registry {
	// Creates the mapping between name and the plugin
	m := make(map[string]plugin.Plugin)
	for _, p := range plugins {
		if _, ok := m[p.Name()]; ok {
			panic("Duplicate plugin found with name: " + p.Name())
		}

		m[p.Name()] = p
	}

	log.Println("Identified plugins: ", m)
	// Create the registry with the mapping.
	return defaultRegistry{m}
}

type defaultRegistry struct {
	pluginsMap map[string]plugin.Plugin
}

func (r defaultRegistry) Run(pluginName string) (interface{}, error) {
	if p := r.pluginsMap[pluginName]; p == nil {
		return nil, ErrPluginNotFound
	} else {
		log.Println("Executing plugin: ", pluginName)
		return p.Exec()
	}
}

func (r defaultRegistry) PluginNames() []string {
	var names []string
	for n, _ := range r.pluginsMap {
		names = append(names, n)
	}
	return names
}
