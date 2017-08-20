package main

import (
	"github.com/ravichaturvedi/go-monitor/registry"
	"github.com/ravichaturvedi/go-monitor/plugin"
	"github.com/ravichaturvedi/go-monitor/server"
	"log"
	"github.com/ravichaturvedi/go-monitor/handler"
)


func main() {
	hwp := plugin.New("helloworld", func() (interface{}, error) {
		return "Hello World!", nil
	})

	r := registry.New(hwp)
	s := server.New(handler.New(r))

	if err := s.Serve(); err != nil {
		log.Fatal("Error while starting server: ", err)
	}
}

