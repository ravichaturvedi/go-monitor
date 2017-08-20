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