package handler

import (
	"net/http"
	"encoding/json"
	"github.com/ravichaturvedi/go-monitor/registry"
	"fmt"
)


func New(r registry.Registry) http.Handler {
	return requestHandler{r}
}

type requestHandler struct {
	r registry.Registry
}

func (h requestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.RequestURI == "/" {
		h.rootHandler(writer, request)
		return
	}
}

func (h requestHandler) rootHandler(writer http.ResponseWriter, request *http.Request) {
	if data, err := json.Marshal(h.r.PluginNames()); err == nil {
		writer.Write(data)
	} else {
		writer.WriteHeader(500)
		writer.Write([]byte(fmt.Sprintf("%v", err)))
	}
}