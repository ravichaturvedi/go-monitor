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
	hwp := plugin.New(func() (interface{}, error) {
		return "Hello World!", nil
	})

	r := registry.New(map[string]plugin.Plugin{"helloworld": hwp})
	r = scheduler.New(r, map[string]time.Duration{"helloworld": time.Second})

	s := server.New(handler.New(r))

	if err := s.Serve(); err != nil {
		log.Fatal("Error while starting server: ", err)
	}
}

