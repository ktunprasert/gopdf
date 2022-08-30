package routes

import (
	"html/template"

	"github.com/joofjang/numgothai"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type templateVariable struct {
    Value interface{}
}

func (v *templateVariable) Set(value interface{}) string {
    v.Value = value
    return ""
}

func newTemplateVariable(initValue interface{}) *templateVariable{
    return &templateVariable{initValue}
}

var funcMap = template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
	"sub": func(a, b int) int {
		return a - b
	},
	"mul": func(a, b int) int {
		return a * b
	},
	"div": func(a, b int) int {
		return a / b
	},
	"printDecAsFloat": func(a int) string {
        p := message.NewPrinter(language.English)
        return p.Sprintf("%.2f", float64(a) / 100)
	},
    "var": newTemplateVariable,
    "percentage": func(a int, percent float64) int {
        val := float64(a) * percent
        return int(val)
    },
    "bahtThaiText": numgothai.IntBaht,
}
