package gohttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

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

// httpGetCustom 自动移 http 请求，并设置 3 秒超时时间。
func httpGetCustom(url string) (resp *http.Response, err error) {
	cli := http.Client{
		Timeout: 3 * time.Second,
	}
	return cli.Get(url)
}
