package types

import (
	"encoding/json"
	"strconv"
)

type Categories map[int64]string

func (c *Categories) UnmarshalJSON(b []byte) (err error) {
	var m map[string]string
	if err = json.Unmarshal(b, &m); err != nil {
		return
	}
	if len(m) == 0 {
		return
	}
	var i int64
	*c = make(Categories, len(m))
	for k, v := range m {
		if i, err = strconv.ParseInt(k, 10, 64); err != nil {
			break
		}
		(*c)[i] = v
	}
	return
}
