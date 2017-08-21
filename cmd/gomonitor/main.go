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
package main

import (
	"time"
	"context"
	"math/rand"
	"github.com/ravichaturvedi/go-monitor/plugin"
	"github.com/ravichaturvedi/go-monitor/monitor"
)


func main() {
	p := plugins()
	monitor.New(p, pluginsDuration(p)).Start(context.Background())
}



func plugins() map[string]plugin.Plugin {
	return map[string]plugin.Plugin {
		"hello": plugin.NewHelloWorld(),
		"random": plugin.NewRandomFailure(),
		"google": plugin.NewURLStatus("https://www.google.com"),
		"facebook": plugin.NewURLStatus("https://www.facebook.com"),
	}
}


func pluginsDuration(plugins map[string]plugin.Plugin) map[string]time.Duration {
	m := map[string]time.Duration{}
	for name, _ := range plugins {
		m[name] = time.Duration(rand.Intn(3) + 1) * time.Second
	}
	return m
}