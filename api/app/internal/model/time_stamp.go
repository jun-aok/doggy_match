package model

import "time"

type TimeStamp time.Time

func (t *TimeStamp) UnmarshalParam(src string) error {
	ts, err := time.Parse(time.RFC3339, src)
	*t = TimeStamp(ts)
	return err
}
