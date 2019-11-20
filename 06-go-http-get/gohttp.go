package gohttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tangx/go-examples/utils"
)

func httpGet(URL string) {

	if !strings.HasPrefix(URL, "http") {
		URL = "http://" + URL
	}
	resp, err := http.Get(URL)
	utils.PanicError(err)
	defer resp.Body.Close()

	fmt.Println("StatusCode: ", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	utils.PanicError(err)

	fmt.Println("http resp Body: ", string(body))
}
