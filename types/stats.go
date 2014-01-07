package types

import (
	"encoding/json"
	"strconv"
)

type Stats struct {
	Confidence float64
}

func (s *Stats) UnmarshalJSON(b []byte) (err error) {
	var m map[string]string
	if err = json.Unmarshal(b, &m); err != nil {
		return
	}
	if _, ok := m["confidence"]; ok {
		var f float64
		if f, err = strconv.ParseFloat(m["confidence"], 64); err != nil {
			return
		}
		s.Confidence = f
	}
	return
}

