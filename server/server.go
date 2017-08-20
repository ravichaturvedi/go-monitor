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
)


// Server is interface to expose the registry to the outside world.
type Server interface {
	Serve() error
}

func New(h http.Handler) Server {
	return httpServer{h}
}


type httpServer struct {
	h http.Handler
}


func (s httpServer) Serve() error {
	// Register the handler to respond to the queries.
	http.Handle("/", s.h)

	log.Println(fmt.Sprintf("Starting server: http://0.0.0.0:1234"))

	// Start the server to listen on all the interfaces on port `1234`
	return http.ListenAndServe(":1234", nil)
}