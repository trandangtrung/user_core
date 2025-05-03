package template

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/os/gview"
)

func toUpper(str string) string {
	return strings.ToUpper(str)
}

type ITemplate interface {
	Get(path string, nameFile string, data map[string]interface{}) (string, error)
}

type Template struct {
}

func NewTemplate() ITemplate {
	return &Template{}
}

func (t *Template) Get(path string, nameFile string, data map[string]interface{}) (string, error) {
	view := gview.New()

	view.SetPath(path)
	// view.SetPath("resource/template")

	// view.BindFuncMap(g.MapStrAny{
	// 	"toUpper": toUpper,
	// })

	tmpl, err := view.Parse(context.Background(), nameFile, data)

	if err != nil {

		return "", err
	}

	return tmpl, err
}
