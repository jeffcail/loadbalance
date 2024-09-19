package random

import (
	"errors"
	"math/rand"
)

type RandomBalancing struct {
	curlIndex int
	rss       []string
}

func (r *RandomBalancing) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	r.rss = append(r.rss, addr)
	return nil
}

func (r *RandomBalancing) Next() string {
	if len(r.rss) == 0 {
		return ""
	}
	r.curlIndex = rand.Intn(len(r.rss))
	return r.rss[r.curlIndex]
}

func (r *RandomBalancing) Get(key string) (string, error) {
	return r.Next(), nil
}
