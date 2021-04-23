package timeutil

import "time"

func Null() time.Time {
	return time.Unix(0, 0)
}
