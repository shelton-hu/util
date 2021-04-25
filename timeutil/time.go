package timeutil

import "time"

const (
	DateTimeLayout = "2006-01-02 15:04:05"
	DateLayout     = "2006-01-02"
	TimeLayout     = "15:04:05"

	RFC3339 = time.RFC3339
)

func Null() time.Time {
	return time.Unix(0, 0)
}

func NowStr() string {
	return time.Now().Format(DateTimeLayout)
}
