package timeinfo

import (
	"testing"
	"time"
)

func Test_timeFormat(t *testing.T) {
	now := time.Now()
	time2String(now)

}

func Test_string2Time(t *testing.T) {
	string2Time()
}

func Test_TimeDelta(t *testing.T) {
	timeDelta(time.Now())
}
