/*
 * Copyright 2017 The go-monitor AUTHORS.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
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
