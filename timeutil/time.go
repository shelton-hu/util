package timeutil

import (
	"errors"
	"time"
)

var (
	ErrStratTimeIsGreaterThanEndTime = errors.New("开始时间大于结束时间")
)

func EmptyTime() time.Time {
	return time.Unix(0, 0)
}

func FullTime() time.Time {
	return RFC3339Time("2999-01-01T00:00:00Z")
}

func RFC3339Time(s interface{}) time.Time {
	var t time.Time
	var err error

	switch r := s.(type) {
	case string:
		t, err = String2Time(time.RFC3339, r)
		if err != nil {
			return EmptyTime()
		}
	case int64:
		t = Timestamp2Time(r, 0)
	default:
		return EmptyTime()
	}
	return t
}

func String2Time(layout, str string) (time.Time, error) {
	return time.Parse(layout, str)
}

func Time2String(t time.Time, layout string) string {
	return t.Format(layout)
}

func String2Timestamp(layout, str string, isContainNano bool) (int64, error) {
	t, err := time.Parse(layout, str)
	if err != nil {
		return 0, err
	}
	if isContainNano {
		return t.UnixNano(), nil
	}
	return t.Unix(), nil
}

func Timestamp2String(sec, nsec int64, layout string) string {
	return time.Unix(sec, nsec).Format(layout)
}

func Timestamp2Time(sec, nsec int64) time.Time {
	return time.Unix(sec, nsec)
}

func Time2Timestamp(t time.Time, isContainNano bool) int64 {
	if isContainNano {
		return t.Unix()
	}
	return t.UnixNano()
}

func ValidateStartTimeAndEndTime(startTime string, endTime string) error {
	if startTime != "" && endTime != "" {
		if startTime > endTime {
			return ErrStratTimeIsGreaterThanEndTime
		}
	}
	return nil
}
