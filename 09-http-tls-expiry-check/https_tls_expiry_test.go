package httptlsexpiry

import "testing"

const (
	URL = "https://www.baidu.com"
)

func Test_httpsExpire(t *testing.T) {
	httpGet(URL)
}
