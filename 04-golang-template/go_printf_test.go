package gotemplate

// https://www.cnblogs.com/yinzhengjie/p/7680829.html

import (
	"encoding/json"
	"fmt"
	"testing"

	"gopkg.in/yaml.v2"
)

func Test_FormatPrintf(t *testing.T) {
	// int
	FormatInt(10)
	// float
	FormatFloat(321.1356)
}

/*
	整数
*/

func FormatInt(i int32) {
	fmt.Printf("// %d\n", i)    // 10
	fmt.Printf("// %10d\n", i)  //         10
	fmt.Printf("// %-10d\n", i) // 10
}

/*
	浮点小数
*/

func FormatFloat(f float32) {
	fmt.Printf("// %f\n", f)     // 321.135590
	fmt.Printf("// %10.2f\n", f) //     321.14
	fmt.Printf("// %5.9f\n", f)  // 321.135589600
}

/* string or []byte */

func (u *User) Marshal() {
	b, _ := json.Marshal(u)
	fmt.Printf("%v\n", b) // [123 34 110 97 109 101 34 58 34 229 188 160 228 186 140 231 139 151 34 44 34 97 103 101 34 58 50 48 44 34 98 105 108 108 105 110 103 34 58 49 50 51 52 51 46 50 52 55 125]
	fmt.Printf("%s\n", b) // {"name":"张二狗","age":20,"billing":12343.247}
	fmt.Printf("%x\n", b) // 7b226e616d65223a22e5bca0e4ba8ce78b97222c22616765223a32302c2262696c6c696e67223a31323334332e3234377d
	fmt.Println(string(b))

	yamlB, _ := yaml.Marshal(u)
	fmt.Printf("%s", yamlB)
}

func Test_UserMarshal(t *testing.T) {
	p.Marshal()
}

func Test_Colorful(t *testing.T) {
	
	fmt.Println("")

	// 前景 背景 颜色
	// ---------------------------------------
	// 30  40  黑色
	// 31  41  红色
	// 32  42  绿色
	// 33  43  黄色
	// 34  44  蓝色
	// 35  45  紫红色
	// 36  46  青蓝色
	// 37  47  白色
	//
	// 代码 意义
	// -------------------------
	//  0  终端默认设置
	//  1  高亮显示
	//  4  使用下划线
	//  5  闪烁
	//  7  反白显示
	//  8  不可见

	for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
		for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
			for d := range []int{0, 1, 4, 5, 7, 8} { // 显示方式 = 0,1,4,5,7,8
				fmt.Printf(" %c[%d;%d;%dm%s(f=%d,b=%d,d=%d)%c[0m \n", 0x1B, d, b, f, "", f, b, d, 0x1B)
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}
