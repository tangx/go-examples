package timeinfo

import (
	"fmt"
	"time"
)

const (
	FORMAT = `2006-01-02T15:04:05Z`
)

func printTime(t time.Time) {
	fmt.Println("->", t.Format(FORMAT))
}

func time2String(t time.Time) {

	fmt.Println("time now =", t)

	fmt.Println("time.RFC1123 : ", t.Format(time.RFC1123))
	fmt.Println("time.RFC822Z : ", t.Format(time.RFC822Z))
	fmt.Println("time.RFC3339 : ", t.Format(time.RFC3339))

	FORMAT := `2006-01-02T15:04:05Z`
	fmt.Println(FORMAT, ":", t.Format(FORMAT))

	text := t.AppendFormat([]byte("Time: "), time.RFC1123)
	fmt.Println(string(text))

	fmt.Println("t1.String() = ", t.String())
	fmt.Printf("\n+++++++++++++++++\ns")
}

func string2Time() {

	t1 := time.Date(2019, time.November, 17, 11, 0, 0, 0, time.UTC)
	printTime(t1)

	timeStr := `2019-11-23T15:23:31Z`
	t, err := time.Parse(FORMAT, timeStr)
	if err != nil {
		panic(err)
	}
	printTime("time.Parse : ", t)
}

func timeZone() {

}

func timeDelta(t time.Time) {

	printTime(t)

	// 3 hours later
	printTime(t.Add(3 * time.Hour))
	// 3 hours before
	printTime(t.Add(-3 * time.Hour))

	// a year later
	printTime(t.AddDate(1, 0, 0))

	before3Hour := time.Now().Add(-3 * time.Hour)
	fmt.Println(t.After(before3Hour))

	// 本地时间
	printTime(t.Local()) //2019-11-15T23:23:52Z
	// UTC 时间
	printTime(t.UTC()) //2019-11-15T15:23:52Z

}
