package service

import (
	"github.com/gogf/gf/v2/errors/gerror"
	mail "github.com/quannv/strongbody-api/utility/gmail"
	rescode "github.com/quannv/strongbody-api/utility/resCode"
	"github.com/quannv/strongbody-api/utility/template"
)

type (
	GmailService interface {
	}

	gmailService struct {
		gmail    mail.EmailSender
		template template.ITemplate
	}
)

func NewGmailService(gmail mail.EmailSender, template template.ITemplate) GmailService {
	return &gmailService{
		gmail:    gmail,
		template: template,
	}
}

func (g *gmailService) Welcome(data map[string]interface{}, to []string, attachFiles []string) error {
	tmpl, err := g.template.Get("/resouce/template/welcome", "index.html", data)

	if err != nil {
		return gerror.NewCode(rescode.InternalError, "get template error")

	}

	subject := "A test email"
	content := tmpl

	err = g.gmail.SendEmail(subject, content, to, nil, nil, attachFiles)

	if err != nil {
		return gerror.NewCode(rescode.InternalError, "send email error")
	}

	return nil
}
