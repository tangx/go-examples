package googlevalues

import (
	"fmt"
	"testing"

	"github.com/tangx/go-querystring/query"
)

type Person struct {
	Name    string   `url:"name,omitempty"`
	Age     int      `url:"age,omitempty"`
	Address []string `url:"address,omitempty,dotnumbered,numbered1"`
}

var user = Person{
	Name:    "zhangsan",
	Age:     10,
	Address: []string{"sichuan", "chengdu"},
}

func Test_GoogleQueryValues(t *testing.T) {

	values, _ := query.Values(user)

	fmt.Println(values.Encode()) // address.1=sichuan&address.2=chengdu&age=10&name=zhangsan

}
