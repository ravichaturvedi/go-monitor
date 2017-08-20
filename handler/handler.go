package handler

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/ravichaturvedi/go-monitor/registry"
	"strings"
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
	writeTo(writer, h.r.PluginNames())
}


// pluginHandler provide response to the specific plugin.
func (h requestHandler) pluginHandler(writer http.ResponseWriter, request *http.Request) {
	writeTo(writer, h.r.Run(strings.TrimPrefix(request.RequestURI, "/")))
}


func writeTo(writer http.ResponseWriter, v interface{}) {
	if data, err := json.Marshal(v); err == nil {
		writer.Write(data)
	} else {
		writer.WriteHeader(500)
		writer.Write([]byte(fmt.Sprintf("%v", err)))
	}
}