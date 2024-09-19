package weighted

import (
	"errors"
	"strconv"
)

type WeightBalance struct {
	curIndex int
	rss      []*WeightedNode
	rsw      []int
}

type WeightedNode struct {
	addr            string
	weight          int
	currentWeight   int
	effectiveWeight int
}

func (r *WeightBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("param len need 2")
	}
	parInt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return nil
	}
	node := &WeightedNode{addr: params[0], weight: int(parInt)}
	node.effectiveWeight = node.weight
	r.rss = append(r.rss, node)
	return nil
}

func (r *WeightBalance) Next() string {
	total := 0
	var best *WeightedNode
	for i := 0; i < len(r.rss); i++ {
		w := r.rss[i]
		total += w.effectiveWeight
		w.currentWeight += w.effectiveWeight

		if w.effectiveWeight < w.weight {
			w.effectiveWeight++
		}
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
	}
	if best == nil {
		return ""
	}

	best.currentWeight -= total
	return best.addr
}
