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
package server

import (
	"net/http"
	"log"
	"fmt"
	"time"
	"context"
)


// Server is interface to expose the registry to the outside world.
type Server interface {
	Serve() error
	Shutdown() error
}

func New(h http.Handler) Server {
	return &httpServer{h, nil}
}


type httpServer struct {
	h http.Handler
	hs *http.Server
}


func (s *httpServer) Serve() error {
	hs := &http.Server{
		Addr:           ":1234",
		Handler:        s.h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	log.Println(fmt.Sprintf("Starting server: http://0.0.0.0:1234"))
	return hs.ListenAndServe()
}

func (s *httpServer) Shutdown() error {
	return s.hs.Shutdown(context.TODO())
}