package main

import (
	"log"
	"time"
	"github.com/ravichaturvedi/go-monitor/registry"
	"github.com/ravichaturvedi/go-monitor/plugin"
	"github.com/ravichaturvedi/go-monitor/server"
	"github.com/ravichaturvedi/go-monitor/handler"
	"github.com/ravichaturvedi/go-monitor/scheduler"
)


func main() {
	r := registry.New(map[string]plugin.Plugin{"hello": plugin.NewHelloWorld()})
	r = scheduler.New(r, map[string]time.Duration{"hello": time.Second})

	s := server.New(handler.New(r))

	if err := s.Serve(); err != nil {
		log.Fatal("Error while starting server: ", err)
	}
}

