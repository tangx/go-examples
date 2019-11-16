package gotemplate

// User info
type User struct {
	Name      string   `json:"name,omitempty" yaml:"name,omitempty"`
	Age       int      `json:"age,omitempty" yaml:"age,omitempty"`
	Cellphone string   `json:"cellphone,omitempty" yaml:"cellphone,omitempty"`
	Address   []string `json:"address,omitempty" yaml:"address,omitempty"`
	Billing   float32  `json:"billing,omitempty" yaml:"billing,omitempty"`
}

var p User = User{
	Name:    "张二狗",
	Age:     20,
	Billing: 12343.2472,
	Address: []string{
		"chengdu",
		"shanghai",
		"beijing",
	},
}
