package types

import (
	"encoding/json"
	"fmt"
)

type Position struct {
	X int
	Y int
}

func (p *Position) UnmarshalJSON(b []byte) (err error) {
	var a []int
	if err = json.Unmarshal(b, &a); err != nil {
		return
	}
	if l := len(a); l != 2 {
		err = fmt.Errorf("json: position: incorrect length %d", l)
		return
	}
	p.X = a[0]
	p.Y = a[1]
	return
}

