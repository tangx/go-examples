package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type weather struct{}

func (w weather) Get() {

	b := w.MustReadFromURLorFile("w.html")
	fmt.Println(string(b))

}

func (w weather) GetURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (w weather) MustReadFromURLorFile(path string) []byte {

	if strings.HasPrefix(path, "http") {
		b, err := w.GetURL(path)
		if err != nil {
			panic(err)
		}
		return b
	} else {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		return b
	}
}
