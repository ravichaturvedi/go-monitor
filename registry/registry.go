package registry

import (
	"log"
	"errors"
	"github.com/ravichaturvedi/go-monitor/plugin"
	"sync"
)

var ErrPluginNotFound = errors.New("Plugin not present in registry.")


// Registry holds the available plugins and expose operations to be performed on those plugins.
type Registry interface {

	// Run the specified plugin names and return back the result.
	Run(pluginNames ...string) []plugin.Result

	// List the available plugin names
	PluginNames() []string
}


func New(plugins map[string]plugin.Plugin) (r Registry) {
	// Identify the installed plugins from the returned registry.
	defer func() {
		log.Println("Installed plugins: ", r.PluginNames())
	}()

	// Create the registry with the mapping.
	return defaultRegistry{plugins}
}

type defaultRegistry struct {
	pluginsMap map[string]plugin.Plugin
}

func (r defaultRegistry) Run(pluginNames ...string) []plugin.Result {
	res := make([]plugin.Result, len(pluginNames))
	var wg sync.WaitGroup
	wg.Add(len(pluginNames))

	// Forking go-routine for all the plugins.
	for i, pluginName := range pluginNames {
		go func(i int, pluginName string) {
			res[i] = r.run(pluginName)
			wg.Done()
		}(i, pluginName)
	}

	// waiting for all the results to be returned.
	wg.Wait()
	return res
}

func (r defaultRegistry) run(pluginName string) plugin.Result {
	if p := r.pluginsMap[pluginName]; p == nil {
		return plugin.Result{nil, ErrPluginNotFound}
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
