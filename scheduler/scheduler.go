package scheduler

import (
	"time"
	"sync"
	"github.com/ravichaturvedi/go-monitor/registry"
	"github.com/ravichaturvedi/go-monitor/plugin"
	"log"
)

// Scheduler schedule the plugins to run after every provided duration and exposes the Registry interface.
// This make it a in-place replacement of Registry in case we want to use scheduled plugins and can use the cached data.
//
// If the scheduling information is not found for the plugin name then delegate the call to underlying registry.
type Scheduler interface {
	registry.Registry
}


func New(r registry.Registry, m map[string] time.Duration) Scheduler {
	s := &durationScheduler{r:r, results: make(map[string] plugin.Result)}
	s.start(m)
	return s
}


type durationScheduler struct {
	r registry.Registry
	results map[string] plugin.Result
	sync.Mutex
}

func (s durationScheduler) Run(pluginNames ...string) []plugin.Result {
	var wg sync.WaitGroup
	res := make([]plugin.Result, len(pluginNames))

	// Try to get the cache plugin Result populated during scheduled execution
	// Otherwise, get the result from the underlying registry.
	for i, pluginName := range pluginNames {
		if r, ok := s.results[pluginName]; ok {
			log.Println("Returning cached result for plugin: " + pluginName)
			res[i] = r
		} else {
			wg.Add(1)
			go func(i int, pluginName string) {
				res[i] = s.r.Run(pluginName)[0]
				wg.Done()
			}(i, pluginName)
		}
	}

	wg.Wait()
	return res
}

func (s durationScheduler) PluginNames() []string {
	return s.r.PluginNames()
}

func (s *durationScheduler) start(m map[string] time.Duration) {
	defer func() {
		log.Println("Started scheduler with: ", m)
	}()

	var wg sync.WaitGroup
	for name, duration := range m {
		wg.Add(1)
		go func(name string, duration time.Duration) {
			s.execPlugin(name)
			wg.Done()
			s.schedulePlugin(name, duration)
		}(name, duration)
	}

	// waiting for all the plugins to populate first result.
	wg.Wait()
}

func (s *durationScheduler) schedulePlugin(name string, d time.Duration) {
	for range time.NewTicker(d).C {
		s.execPlugin(name)
	}
}

func (s *durationScheduler) execPlugin(name string) {
	r := s.r.Run(name)

	// Store the value in result map
	s.Lock()
	s.results[name] = r[0]
	s.Unlock()
}