package rotation

import "errors"

type RoundRotationBalance struct {
	curIndex int
	rss      []string
	//conf LoadBalanceConf
}

func (r *RoundRotationBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}

	addr := params[0]
	r.rss = append(r.rss, addr)
	return nil
}

func (r *RoundRotationBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}

	lens := len(r.rss)
	if r.curIndex >= lens {
		r.curIndex = 0
	}
	curAddr := r.rss[r.curIndex]
	r.curIndex = (r.curIndex + 1) % lens
	return curAddr
}
