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

func (g *gmailService) Welcome(userName string, to []string, attachFiles []string) error {
	data := map[string]interface{}{
		"UserName":       userName,
		"SupportContact": "examle@strongbody.ai",
	}
	tmpl, err := g.template.Get("/resouce/template/welcome", "index.html", data)

	if err != nil {
		return gerror.NewCode(rescode.InternalError, "get template error")

	}

	subject := "Welcome to Strongbody! Start Your Health Journey Today!"
	content := tmpl

	err = g.gmail.SendEmail(subject, content, to, nil, nil, attachFiles)

	if err != nil {
		return gerror.NewCode(rescode.InternalError, "send email error")
	}

	return nil
}

func (g *gmailService) CodeOtp(userName string, code int, to []string, attachFiles []string) error {
	data := map[string]interface{}{
		"UserName":   userName,
		"VerifyCode": code,
	}
	tmpl, err := g.template.Get("/resouce/template/code-otp", "index.html", data)

	if err != nil {
		return gerror.NewCode(rescode.InternalError, "get template error")

	}

	subject := "Verify Your Email to Complete Your StrongBody SignUp!"
	content := tmpl

	err = g.gmail.SendEmail(subject, content, to, nil, nil, attachFiles)

	if err != nil {
		return gerror.NewCode(rescode.InternalError, "send email error")
	}

	return nil
}
