package types

import (
	"encoding/json"
	"time"
)

const (
	// Modified RFC1123 format
	mRFC1123 = "Mon, 2 Jan 2006 15:04:05 MST"
)

type Time time.Time

func (t Time) Time() time.Time {
	return time.Time(t)
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	var tt time.Time
	if err = json.Unmarshal(b, &tt); err == nil {
		*t = Time(tt)
		return // correct encoding
	}
	// Try decoding as a string
	var s string
	if err = json.Unmarshal(b, &s); err != nil {
		return
	}
	tt, err = time.Parse(mRFC1123, s)
	*t = Time(tt)
	return
}
