package pool_obj

import (
	"fmt"
	"testing"
)

func TestGetPool(t *testing.T) {
	p:= NewPool(10)
	o, err:=p.GetObj()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(len(p.pool_ch))
	fmt.Printf("%v\n", o)
	p.ReleaseObj(o)
	t.Log(len(p.pool_ch))
	p.ReleaseObj(o)
}