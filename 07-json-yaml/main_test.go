package jsonyaml

import "testing"

func Test_Marshal(t *testing.T) {
	Marshal()
}

func Test_Yaml2Json(t *testing.T) {
	Yaml2Json()
}

func Test_Json2Yaml(t *testing.T) {
	Json2Yaml()
}
func Test_Yaml2Yaml(t *testing.T) {
	Yaml2Yaml()
}

func Test_JsonMinify(t *testing.T) {
	JsonMinify("config.json")
}

func Test_JsonPretty(t *testing.T) {
	JsonPretty("config.json")
}
