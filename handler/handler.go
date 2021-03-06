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
package handler

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/ravichaturvedi/go-monitor/registry"
	"strings"
	"github.com/ravichaturvedi/go-monitor/plugin"
)


func New(r registry.Registry) http.Handler {
	return requestHandler{r}
}

type requestHandler struct {
	r registry.Registry
}

func (h requestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Only accepting get methods from the http requests.
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if request.RequestURI == "/" {
		h.rootHandler(writer, request)
		return
	}

	h.pluginHandler(writer, request)
}


// rootHandler list result of all the plugins
func (h requestHandler) rootHandler(writer http.ResponseWriter, request *http.Request) {
	pluginNames := h.r.PluginNames()
	results := h.r.Run(pluginNames...)

	// Create mapping between, plugin name and result.
	m := make(map[string]plugin.Result)
	for i, pluginName := range pluginNames {
		m[pluginName] = results[i]
	}
	writeTo(writer, m)
}


// pluginHandler provide response to the specific plugin.
func (h requestHandler) pluginHandler(writer http.ResponseWriter, request *http.Request) {
	writeTo(writer, h.r.Run(strings.TrimPrefix(request.RequestURI, "/"))[0])
}


func writeTo(writer http.ResponseWriter, v interface{}) {
	if data, err := json.Marshal(v); err == nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.Write(data)
	} else {
		writer.WriteHeader(500)
		writer.Write([]byte(fmt.Sprintf("%v", err)))
	}
}