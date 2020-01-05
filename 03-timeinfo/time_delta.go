package timeinfo

import (
	"fmt"
	"time"
)

// timeDelta2 比较两个时间的时间差
func timeDelta2() {

	t1 := time.Now()

	time.Sleep(time.Second * 1)
	t2 := time.Now()

	d1 := t1.Sub(t2)
	fmt.Println(d1.Seconds())

	d2 := t2.Sub(t1)
	fmt.Println(d2.Seconds())

}
