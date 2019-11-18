package gotemplate

// https://www.cnblogs.com/f-ck-need-u/p/10053124.html

import (
	"os"
	"testing"
	"text/template"
)

func Test_gotemplateSimple(t *testing.T) {
	p := map[string]string{
		"User":   "王二狗",
		"Age":    "10",
		"Gender": "Boy",
	}
	// 模板定义
	tepl := `My name is {{ .User }}, 
	I'm {{ .Age }} years old,
	I'm a {{ .Gender }}.
	`

	// 解析模板
	tmpl, _ := template.New("test").Parse(tepl)

	// 数据驱动模板
	// data := "    "
	_ = tmpl.Execute(os.Stdout, p)
}

func Test_ifTemplate(t *testing.T) {
	p := []string{
		"zhangsan",
		"lisi",
		"wangwu",
		"zhaoerwa",
	}

	// range . 中的 `.` 的作用域是  p
	// {{ . }} 中的 `.` 的作用域是 []string 的值
	tepl := `{{ range . }} {{ . }} {{ end }}`
	tmpl, _ := template.New("test").Parse(tepl)
	_ = tmpl.Execute(os.Stdout, p)

}

func Test_rangeListMap(t *testing.T) {

	// ex2s
	u := []map[string]string{
		{
			"User":   "王二狗",
			"Age":    "10",
			"Gender": "Boy",
		},
		{
			"User":   "张小丫",
			"Age":    "20",
			"Gender": "Girl",
		},
	}

	tepl2 := `
	{{ range . }}
	{{ .User }}: a {{ .Age }} year {{ .Gender }}
	{{ end }}
	`
	tmpl2, _ := template.New("test").Parse(tepl2)
	tmpl2.Execute(os.Stdout, u)
}

func Test_templateMapMap(t *testing.T) {
	// ex2s
	u := map[string]map[string]string{
		"boy": {
			"User":   "王二狗",
			"Age":    "10",
			"Gender": "Boy",
		},
		"girl": {
			"User":   "张小丫",
			"Age":    "20",
			"Gender": "Girl",
		},
	}

	tepl2 := `
	{{ .boy.User }}
	{{ .girl.User }}
	`
	tmpl2, _ := template.New("test").Parse(tepl2)
	tmpl2.Execute(os.Stdout, u)
}

func Test_rangeMapList(t *testing.T) {
	// ex4
	u := map[string][]string{
		"student": []string{"  王二狗  ", "  张二丫  ", "  李大锤  "},
		"teacher": []string{"慕容大脸", "诸葛大胆", "欧阳大腿"},
	}

	tepl2 := `
	{{ range .teacher -}}
    {{ . }}  
	{{ end }}
	`
	tmpl2, _ := template.New("test").Parse(tepl2)
	tmpl2.Execute(os.Stdout, u)
}
