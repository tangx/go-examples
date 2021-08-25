package timeinfo

import (
	"fmt"
	"testing"
	"time"
)

func Test_Compare(t *testing.T) {

	t1, _ := time.Parse("2006-01-02T15:04:05Z", "2021-09-20T10:12:30")
	t2, _ := time.Parse("2006-01-02T15:04:05Z", "2021-08-20T10:12:30")

	fmt.Println(t1.After(t2))
	fmt.Println(t1.Equal(t1))
}
