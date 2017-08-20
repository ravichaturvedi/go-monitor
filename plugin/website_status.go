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

import (
	"net/http"
	"io/ioutil"
	"io"
)

func NewURLStatus(url string) Plugin {
	return New(func() (interface{}, error) {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()
		io.Copy(ioutil.Discard, res.Body)
		return res.Status, nil
	})
}
