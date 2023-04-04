package model

import (
	"fmt"
	"time"
)

type Date time.Time

func (t *Date) UnmarshalParam(src string) error {
	ts, err := time.Parse("2006-01-02", src)
	if err != nil {
		return fmt.Errorf("誕生日は日付形式で入力してください")
	}
	*t = Date(ts)
	return nil
}
