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
