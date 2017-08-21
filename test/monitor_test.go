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
