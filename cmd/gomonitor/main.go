package main

import (
	"time"
	"math/rand"
	"github.com/ravichaturvedi/go-monitor/plugin"
	"github.com/ravichaturvedi/go-monitor/monitor"
)


func main() {
	p := plugins()
	monitor.New(p, pluginsDuration(p)).Start()
}



func plugins() map[string]plugin.Plugin {
	return map[string]plugin.Plugin {
		"hello": plugin.NewHelloWorld(),
		"random": plugin.NewRandomFailure(),
		"google": plugin.NewWebsiteStatus("https://www.google.com"),
		"facebook": plugin.NewWebsiteStatus("https://www.facebook.com"),
	}
}


func pluginsDuration(plugins map[string]plugin.Plugin) map[string]time.Duration {
	m := map[string]time.Duration{}
	for name, _ := range plugins {
		m[name] = time.Duration(rand.Intn(3) + 1) * time.Second
	}
	return m
}