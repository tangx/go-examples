package timeinfo

import (
	"fmt"
	"testing"
	"time"
)

func Test_TimeAttr(t *testing.T) {
	now := time.Now()

	b, _ := now.MarshalJSON()
	fmt.Printf("%s\n", b)

	b2, _ := now.MarshalText()
	fmt.Printf("%s\n", b2)

	now.Year()    // 2021
	now.YearDay() // 237
	now.Month()   // August int(xxx)
	now.Day()     // 25
	now.Date()    // 2021 August 25

}
