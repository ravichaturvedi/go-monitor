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
package plugin

import "encoding/json"

// Result is the response type of plugin execution.
type Result struct {
	Val interface{}
	Err error
}

func (r Result) MarshalJSON() ([]byte, error) {
	if r.Err == nil {
		return json.Marshal(map[string]interface{}{"msg": r.Val})
	}

	return json.Marshal(map[string]interface{}{"msg": r.Val, "err": r.Err.Error()})
}

// Plugin is the core type for extending go-monitor
type Plugin interface {
	Exec() Result
}


// New creates a plugin from the provided function.
func New(execFn func() (interface{}, error)) Plugin {
	return defaultPlugin{execFn}
}


type defaultPlugin struct {
	execFn func() (interface{}, error)
}

func (p defaultPlugin) Exec() Result {
	val, err := p.execFn()
	return Result{val, err}
}
