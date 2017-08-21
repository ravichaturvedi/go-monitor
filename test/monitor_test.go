package test

import (
	"testing"
	"time"
	"context"
	"github.com/ravichaturvedi/go-monitor/monitor"
	"github.com/ravichaturvedi/go-monitor/plugin"
	"net/http"
	"io/ioutil"
	"strings"
)


func TestMonitor(t *testing.T) {
	p := map[string]plugin.Plugin{"hello": plugin.New(func() (interface{}, error) {
		return "Hello World!", nil
	})}

	d := map[string]time.Duration{"hello": time.Second}
	c, f := context.WithTimeout(context.Background(), 2 * time.Second)
	go monitor.New(p, d).Start(c)

	res, err := http.Get("http://localhost:1234/hello")
	if err != nil {
		f()
		t.Error("Got error: " + err.Error())
		return
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		f()
		t.Error("Got error: " + err.Error())
		return
	}

	if !strings.Contains(string(data), "Hello World!") {
		t.Error("Invalid output from monitor.")
	}

	f()
}
