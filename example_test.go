package gcache

import (
	"testing"

	"github.com/carlmjohnson/be"
)

type X struct {
	Time int64
}

func TestCache(t *testing.T) {

	key := "thekey"
	x, err := Cache(key, func() (ret X, err error) {
		return getSomething(10)
	})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	y, err := Cache(key, func() (ret X, err error) {
		return getSomething(20)
	})

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	be.Equal(t, x.Time, y.Time)

}

func getSomething(x int64) (resp X, err error) {
	resp.Time = x
	return
}
