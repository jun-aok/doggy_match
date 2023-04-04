package util

import "time"

// 指定したTimeの00:00:00時点のTimeを取得する
func BeggingOfDay(t time.Time) time.Time {
	ut := t.Unix()
	_, offset := t.Zone()
	return time.Unix((ut/86400)*86400-int64(offset), 0)
}

// 指定したTimeをYYYY-mm-DD形式で取得する
func GetYYYYMMDD(t time.Time) string {
	return t.Format("2006-01-02")
}

func ParseTime(src string) *time.Time {
	t, error := time.Parse(time.RFC3339, src)
	if error != nil {
		return nil
	}
	return &t
}

func ParseDate(src string) *time.Time {
	t, error := time.Parse("2006-01-02", src)
	if error != nil {
		return nil
	}
	return &t
}

// 2006-01-02T15:04:05+09:00形式の時間をTime型に変換する
func JstStringToTime(s string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05+09:00", s)
}
