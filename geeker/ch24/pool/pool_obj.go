package pool_obj

import (
	"errors"
	"time"
)

type Mypool struct {
	pool_ch chan *Myobj
}
type Myobj struct {
}

func NewPool(numofpool int) *Mypool {
	p := new(Mypool)
	p.pool_ch = make(chan *Myobj, numofpool)
	for i := 0; i < numofpool; i++ {
		p.pool_ch <- &Myobj{}
	}
	return p
}

func (p *Mypool) GetObj() (*Myobj, error) {
	select {
	case ret := <-p.pool_ch:
		return ret, nil
	case <-time.After(time.Second * 30):
		return nil, errors.New("time out")
	}
}

func (p *Mypool) ReleaseObj(obj *Myobj) error {
	select {
	case p.pool_ch <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
