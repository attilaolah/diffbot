package types

import (
	"time"
)

type Time time.Time

func (t Time) Time() time.Time {
	return time.Time(t)
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	// TODO: unmarshal! (try the original, the nfall back to a custom format)
	return
}
