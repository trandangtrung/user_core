package mail

import (
	"log"
	"testing"

	"github.com/quannv/strongbody-api/utility/template"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	name := "Leopard"
	fromEmailAddress := ""
	fromEmailPassword := ""

	sender := NewGmailSender(name, fromEmailAddress, fromEmailPassword)

	data := g.Map{
		"UserName":  "Nguyễn Đại Nghĩa",
		"header":    "Test header",
		"container": "Test container 1",
		"footer":    "Test footer",
		"socials": map[string]string{
			"Twitter/X":  "@example-T",
			"Facebook/F": "@example-F",
			"Instagram":  "@examplepics-I",
		},
		"list": []string{
			"Comment",
			"Order",
			"test",
		},
	}

	initTmpl := template.NewTemplate()

	tmpl, err := initTmpl.Get("/resource/template/welcome", "index.html", data)

	if err != nil {
		require.NoError(t, err)

		log.Println(err)
	}

	subject := "A test email"
	content := tmpl
	to := []string{"nghiabeo1605@gmail.com"}
	attachFiles := []string{"../../README.MD"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
