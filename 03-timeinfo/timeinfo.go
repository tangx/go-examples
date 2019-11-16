package timeinfo

import (
	"fmt"
	"time"

	"github.com/tangx/go-examples/utils"
)

const (
	FORMAT = `2006-01-02T15:04:05Z`
)

func printTime(t time.Time) {
	fmt.Println("//", t.Format(FORMAT))
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
	utils.PanicError(err)
	printTime(t)
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

}

func timeZone() {
	now := time.Now()
	printTime(now.Local()) // 2019-11-16T11:22:59Z
	printTime(now.UTC())   // 2019-11-16T03:22:59Z

	// show local location
	curLoc := now.Local().Location()
	fmt.Println(curLoc)
	fmt.Println(now.Location())
	fmt.Println(now.Local().UTC())
	// locTZ := time.LoadLocationFromTZData()

	// locUser := time.LoadLocation("Asia/Shanghai")
	zone, offset := now.Zone()
	fmt.Println("zone:", zone, "\noffset:", offset)

	zone, offset = now.Local().Zone()
	fmt.Println("zone:", zone, "\noffset:", offset)

	zone, offset = now.UTC().Zone()
	fmt.Println("zone:", zone, "\noffset:", offset)

	// // 用户自定义 timezone
	loc, err := time.LoadLocation("America/Los_Angeles")
	utils.PanicError(err)
	fmt.Println("America/Los_Angeles", now.In(loc).Format(FORMAT))
	// more user define timezone
	for _, locStr := range []string{"America/Los_Angeles",
		"Asia/Shanghai",
		"CST",
		"Pacific/Marquesas",
		"America/Boise",
		"Antarctica/Rothera",
		"Asia/Oral"} {
		loc, err := time.LoadLocation("America/Los_Angeles")
		utils.PanicError(err)
		fmt.Println(locStr, now.In(loc).Format(FORMAT))

	}
}
