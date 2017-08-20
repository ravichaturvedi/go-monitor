package plugin

import (
	"math/rand"
	"errors"
)

var ErrFailed = errors.New("Failed")

func NewRandomFailure() Plugin {
	return New(func() (interface{}, error) {
		n := rand.Intn(100)
		if n < 50 {
			return nil, ErrFailed
		}

		return "Success", nil
	})
}
