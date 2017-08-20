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
	hwp := plugin.New("helloworld", func() (interface{}, error) {
		return "Hello World!", nil
	})

	r := registry.New(hwp)
	r = scheduler.New(r, map[string]time.Duration{"helloworld": time.Second})
	s := server.New(handler.New(r))

	if err := s.Serve(); err != nil {
		log.Fatal("Error while starting server: ", err)
	}
}

