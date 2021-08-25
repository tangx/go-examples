package timeinfo

import (
	"fmt"
	"testing"
	"time"
)

func Test_timeFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now.String())
	time2String(now)

}

func Test_string2Time(t *testing.T) {
	string2Time()
}

func Test_TimeDelta(t *testing.T) {
	timeDelta(time.Now())
}

func Test_timeZone(t *testing.T) {
	timeZone()
}
