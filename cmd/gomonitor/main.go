package main

import (
	"log"
	"time"
	"github.com/ravichaturvedi/go-monitor/registry"
	"github.com/ravichaturvedi/go-monitor/plugin"
	"github.com/ravichaturvedi/go-monitor/server"
	"github.com/ravichaturvedi/go-monitor/handler"
	"github.com/ravichaturvedi/go-monitor/scheduler"
	"math/rand"
)


func main() {
	p := plugins()
	r := registry.New(p)
	r = scheduler.New(r, pluginsDuration(p))

	s := server.New(handler.New(r))

	if err := s.Serve(); err != nil {
		log.Fatal("Error while starting server: ", err)
	}
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