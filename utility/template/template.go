package template

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/os/gview"
)

func toUpper(str string) string {
	return strings.ToUpper(str)
}

type TemplateI interface {
	Get(nameFile string, data map[string]interface{}) (string, error)
}

type Template struct {
}

func NewTemplate() TemplateI {
	return &Template{}
}

func (t *Template) Get(nameFile string, data map[string]interface{}) (string, error) {
	view := gview.New()

	view.SetPath("resource/template")

	// view.BindFuncMap(g.MapStrAny{
	// 	"toUpper": toUpper,
	// })

	tmpl, err := view.Parse(context.Background(), nameFile, data)

	if err != nil {

		return "", err
	}

	return tmpl, err
}

// ExampleS
// tmpl,err := template.Get("layout.html",map[string]string{toUpper: "toUpper"})
// sender.SendEmail(subject, tmpl, to, nil, nil, attachFiles)
