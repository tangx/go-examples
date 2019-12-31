package stringsmods

import (
	"fmt"
	"strings"
)

func sCompare() {
	i := strings.Compare("abc", "b")
	fmt.Println(i)
}

func f2() {
	ok := strings.Contains("my name is tangx", "is")
	fmt.Println(ok)

	ok = strings.ContainsAny("my name is tangx", "zhon")
	if ok {
		fmt.Println("ok2")
	}
}

type Person struct {
	Name   string `yaml:"name,omitempty,options2,options3" json:"name,omitempty"`
	Age    int    `yaml:"age,omitempty,options2,options3" json:"age,omitempty"`
	Gender string `yaml:"gender,omitempty,options2,options3" json:"gender,omitempty"`
}
