package jsonyaml

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Person struct {
	Name          string `json:"name,omitempty" yaml:"name,omitempty"`
	Age           int    `json:"age,omitempty" yaml:"age,omitempty"`
	Gender        string `json:"gender,omitempty" yaml:"gender,omitempty"`
	SchoolAddress string `json:"schooladdress,omitempty" yaml:"schooladdress,omitempty"`
}

func Marshal() {

	person := Person{
		Name:   "王小二",
		Age:    10,
		Gender: "男",
	}

	body, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	body, err = yaml.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

}

func Yaml2Json() {
	body, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	// var person Person
	person := Person{}
	err = yaml.Unmarshal(body, &person)
	if err != nil {
		panic(err)
	}

	fmt.Println(person)

	b, _ := json.Marshal(person)
	fmt.Printf("%s", b)
}

func Json2Yaml() {
	body, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	var person Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		panic(err)
	}
	fmt.Println(person.Age)

	b, _ := yaml.Marshal(person)
	fmt.Printf("%s", b)
}

func Yaml2Yaml() {
	body, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	// 注意
	// 1. struct 中的字段为大写是， yaml.Unmarshal 才能将值正确赋值到字段，否则会报错。
	// 2. body(in) 的字段， 首写字母必须为小写，这个是 yaml 的语法结构，否则无法读取。
	person := Person{}
	err = yaml.Unmarshal(body, &person)
	if err != nil {
		panic(err)
	}

	fmt.Println(person)

	b, _ := yaml.Marshal(person)
	fmt.Printf("%s", b)
}

func JsonMinify(filename string) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var p interface{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

	result, _ := json.Marshal(p)
	fmt.Printf("%s", result)

	fobj, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer fobj.Close()

	_, err = fobj.Write(result)
	if err != nil {
		panic(err)
	}
}

func JsonPretty(filename string) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var p interface{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

	result, _ := json.MarshalIndent(p, "", "  ")
	fmt.Printf("%s", result)

	fobj, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer fobj.Close()

	_, err = fobj.Write(result)
	if err != nil {
		panic(err)
	}
}
